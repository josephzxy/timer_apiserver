// Package app is the logical main body of the program.
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

// App defines the interface of an app.
type App interface {
	Run()
}

// New returns a value of an implementation of interface App.
func New(basename string) App {
	return newApp(basename)
}

// app is a concrete implementation of interface App.
// It holds app-level fields like config, cliflags, and the cobra command.
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

// buildCmd initializes the app-level cobra command.
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

// ensureViperValueType ensures viper to store non-string config values with proper types.
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

// bindConfigFromCliFlags binds flags of the cobra command to viper.
func (a *app) bindConfigFromCliFlags() error {
	if err := viper.BindPFlags(a.cmd.Flags()); err != nil {
		msg := "failed to bind configs with cli flags"
		zap.S().Errorw(msg, "err", err)

		return errors.WithMessage(err, msg)
	}

	return nil
}

// bindConfigFromEnv binds env vars to viper.
func (a *app) bindConfigFromEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ToUpper(a.basename))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

// bindConfigFromFile binds a config file to viper.
func (a *app) bindConfigFromFile() error {
	viper.SetConfigFile(a.cfg.Config)
	if err := viper.ReadInConfig(); err != nil {
		msg := "failed to bind configs with config file"
		zap.S().Errorw(msg, "err", err)

		return errors.WithMessage(err, msg)
	}

	return nil
}

// unmarshalConfig unmarshals configs from viper to the app-level config.
func (a *app) unmarshalConfig() error {
	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal from viper to app config"
		zap.S().Errorw(msg, "err", err)

		return errors.WithMessage(err, msg)
	}

	return nil
}

// loadConfig loads (binds and unmarshals) configs to the app-level config
// from various sources like flags, env vars, and files.
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

// Run is the entry point for the app.
func (a *app) Run() {
	defer log.Flush()
	if err := a.cmd.Execute(); err != nil {
		zap.S().Panicw("app failed during running", "err", err)
	}
}

// runCmd is the entry point for the app-level cobra command.
// Note that cli flags and env vars will NOT be read BEFORE calling runCmd.
func (a *app) runCmd(cmd *cobra.Command, args []string) error {
	if err := a.loadConfig(); err != nil {
		return err
	}
	if err := a.run(); err != nil {
		return err
	}

	return nil
}

// run is the entry point for the business logic part of the app where
// database connections are created, server instances are launched, etc.
func (a *app) run() error {
	mysqlRouter, err := mysql.NewRouter(&mysql.Config{
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

	serviceRouter := service.NewRouter(mysqlRouter)
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
