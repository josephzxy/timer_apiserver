package app

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/josephzxy/timer_apiserver/internal/app/cliflags"
	"github.com/josephzxy/timer_apiserver/internal/app/config"
	"github.com/josephzxy/timer_apiserver/internal/grpcserver"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql"
	"github.com/josephzxy/timer_apiserver/internal/restserver"
)

type App struct {
	basename string
	cfg      *config.Config
	cliflags *cliflags.CliFlags
	cmd      *cobra.Command
}

func NewApp(basename string) *App {
	a := &App{
		basename: basename,
		cfg:      config.NewEmptyConfig(),
		cliflags: cliflags.NewCliFlags(),
	}
	a.buildCmd()
	return a
}

func (a *App) buildCmd() {
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
func (a *App) ensureViperValueType() {
	viper.SetDefault("mysql.port", a.cfg.MySQL.Port)
	viper.SetDefault("mysql.parse-time", a.cfg.MySQL.ParseTime)
	viper.SetDefault("mysql.max-idle-conns", a.cfg.MySQL.MaxIdleConns)
	viper.SetDefault("mysql.max-open-conns", a.cfg.MySQL.MaxOpenConns)
	viper.SetDefault("mysql.max-conn-lifetime", a.cfg.MySQL.MaxConnLifetime)
	viper.SetDefault("mysql.log-level", a.cfg.MySQL.LogLevel)

	viper.SetDefault("restserver.insecure-serving.port", a.cfg.RESTServer.InsecureServing.Port)

	viper.SetTypeByDefaultValue(true)
}

func (a *App) bindConfigFromCliFlags() error {
	if err := viper.BindPFlags(a.cmd.Flags()); err != nil {
		msg := "failed to read config from cli flags"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) bindConfigFromEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ToUpper(a.basename))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func (a *App) bindConfigFromFile() error {
	cfgFile := config.CfgFile()
	if cfgFile == "" {
		return fmt.Errorf("config file path must not be empty")
	}
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		msg := "failed to read config file"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) bindConfig() error {
	a.ensureViperValueType()
	if err := a.bindConfigFromCliFlags(); err != nil {
		return err
	}
	a.bindConfigFromEnv()
	if err := a.bindConfigFromFile(); err != nil {
		return err
	}
	return nil
}

func (a *App) loadConfig() error {
	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal config"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		zap.S().Fatal(err)
	}
}

func (a *App) runCmd(cmd *cobra.Command, args []string) error {
	if err := a.bindConfig(); err != nil {
		return err
	}
	if err := a.loadConfig(); err != nil {
		return err
	}
	if err := a.run(); err != nil {
		return err
	}
	return nil
}

func (a *App) run() error {
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
				Host: a.cfg.RESTServer.InsecureServing.Host,
				Port: a.cfg.RESTServer.InsecureServing.Port,
			},
			Mode: a.cfg.RESTServer.Mode,
		},
		serviceRouter,
	)

	var eg errgroup.Group

	eg.Go(func() error {
		if err := restServer.Start(); err != nil {
			msg := "error occured during running rest server"
			zap.S().Errorw(msg, "err", err)
			return errors.WithMessage(err, msg)
		}
		return nil
	})

	grpcServer := grpcserver.New(
		&grpcserver.Config{
			InsecureServing: &grpcserver.InsecureServingConfig{
				Host: "0.0.0.0",
				Port: 8082,
			},
		},
		serviceRouter,
	)
	eg.Go(func() error {
		if err := grpcServer.Start(); err != nil {
			msg := "error occured during running grpc server"
			zap.S().Errorw(msg, "err", err)
			return errors.WithMessage(err, msg)
		}
		return nil
	})
	return eg.Wait()
}
