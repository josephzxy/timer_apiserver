package app

import (
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/josephzxy/timer_apiserver/internal/app/cliflags"
	"github.com/josephzxy/timer_apiserver/internal/app/config"
	"github.com/josephzxy/timer_apiserver/internal/app/gracefulshutdown"
	"github.com/josephzxy/timer_apiserver/internal/grpcserver"
	"github.com/josephzxy/timer_apiserver/internal/pkg/log"
	"github.com/josephzxy/timer_apiserver/internal/pkg/util"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql"
	"github.com/josephzxy/timer_apiserver/internal/restserver"
)

type App interface {
	Run()
}

func New(basename string) App {
	return newApp(basename)
}

type app struct {
	basename string
	cfg      *config.Config
	cliflags cliflags.CliFlags
	cmd      *cobra.Command
}

func newApp(basename string) *app {
	a := &app{
		basename: basename,
		cfg:      config.NewEmptyConfig(),
		cliflags: cliflags.NewCliFlags(),
	}
	a.buildCmd()
	return a
}

func (a *app) buildCmd() {
	a.cmd = &cobra.Command{
		Use:           a.basename,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	a.cmd.RunE = a.runCmd

	for _, fs := range a.cliflags.GetAllFlagSets() {
		a.cmd.Flags().AddFlagSet(fs)
	}
}

// ensureViperValueType ensures viper to store non-string config values with proper types
func (a *app) ensureViperValueType() {
	viper.SetDefault("mysql.port", a.cfg.MySQL.Port)
	viper.SetDefault("mysql.parse-time", a.cfg.MySQL.ParseTime)
	viper.SetDefault("mysql.max-idle-conns", a.cfg.MySQL.MaxIdleConns)
	viper.SetDefault("mysql.max-open-conns", a.cfg.MySQL.MaxOpenConns)
	viper.SetDefault("mysql.max-conn-lifetime", a.cfg.MySQL.MaxConnLifetime)
	viper.SetDefault("mysql.log-level", a.cfg.MySQL.LogLevel)

	viper.SetDefault("restserver.insecure-serving.port", a.cfg.REST.Insecure.Port)
	viper.SetDefault("restserver.middlewares", a.cfg.REST.Middlewares)
	viper.SetDefault("restserver.use-healthz", a.cfg.REST.UseHealthz)

	viper.SetDefault("grpcserver.insecure-serving.port", a.cfg.GRPC.Insecure.Port)

	viper.SetTypeByDefaultValue(true)
}

func (a *app) bindConfigFromCliFlags() error {
	if err := viper.BindPFlags(a.cmd.Flags()); err != nil {
		msg := "failed to bind configs with cli flags"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *app) bindConfigFromEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ToUpper(a.basename))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func (a *app) bindConfigFromFile() error {
	viper.SetConfigFile(a.cfg.Config)
	if err := viper.ReadInConfig(); err != nil {
		msg := "failed to bind configs with config file"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *app) unmarshalConfig() error {
	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal from viper to app config"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *app) loadConfig() error {
	a.ensureViperValueType()
	if err := a.bindConfigFromCliFlags(); err != nil {
		return err
	}
	a.bindConfigFromEnv()
	if err := a.unmarshalConfig(); err != nil {
		return err
	}

	if a.cfg.Config != "" {
		zap.L().Info("config file path set, will read from config file")
		if err := a.bindConfigFromFile(); err != nil {
			return err
		}
		if err := a.unmarshalConfig(); err != nil {
			return err
		}
		return nil
	}
	zap.L().Info("config file path not set, will skip reading from config file")
	return nil
}

func (a *app) Run() {
	defer log.Flush()
	if err := a.cmd.Execute(); err != nil {
		zap.S().Panicw("app failed during running", "err", err)
	}
}

func (a *app) runCmd(cmd *cobra.Command, args []string) error {
	if err := a.loadConfig(); err != nil {
		return err
	}
	if err := a.run(); err != nil {
		return err
	}
	return nil
}

func (a *app) run() error {
	mysqlStoreRouter, err := mysql.NewStoreRouter(&mysql.Config{
		User:            a.cfg.MySQL.User,
		Pwd:             a.cfg.MySQL.Pwd,
		Host:            a.cfg.MySQL.Host,
		Port:            a.cfg.MySQL.Port,
		Database:        a.cfg.MySQL.Database,
		Charset:         a.cfg.MySQL.Charset,
		ParseTime:       a.cfg.MySQL.ParseTime,
		Loc:             a.cfg.MySQL.Loc,
		MaxIdleConns:    a.cfg.MySQL.MaxIdleConns,
		MaxOpenConns:    a.cfg.MySQL.MaxOpenConns,
		MaxConnLifetime: a.cfg.MySQL.MaxConnLifetime,
		LogLevel:        a.cfg.MySQL.LogLevel, // silent
	})
	if err != nil {
		msg := "failed to get mysql store router"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}

	serviceRouter := service.NewRouter(mysqlStoreRouter)
	restServer := restserver.New(
		&restserver.Config{
			InsecureServing: restserver.InsecureServingConfig{
				Host: a.cfg.REST.Insecure.Host,
				Port: a.cfg.REST.Insecure.Port,
			},
			Mode:        a.cfg.REST.Mode,
			Middlewares: a.cfg.REST.Middlewares,
			UseHealthz:  a.cfg.REST.UseHealthz,
		},
		serviceRouter,
	)
	grpcServer := grpcserver.New(
		&grpcserver.Config{
			InsecureServing: &grpcserver.InsecureServingConfig{
				Host: a.cfg.GRPC.Insecure.Host,
				Port: a.cfg.GRPC.Insecure.Port,
			},
		},
		serviceRouter,
	)

	gracefulshutdown.Enable(func() error {
		return util.BatchGoOrErr(
			restServer.Stop,
			func() error {
				grpcServer.Stop()
				return nil
			},
		)
	})

	return util.BatchGoOrErr(
		restServer.Start,
		grpcServer.Start,
	)
}
