package app

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/josephzxy/timer_apiserver/internal/app/cliflags"
	"github.com/josephzxy/timer_apiserver/internal/app/config"
	"github.com/josephzxy/timer_apiserver/internal/app/gracefulshutdown"
	"github.com/josephzxy/timer_apiserver/internal/grpcserver"
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

	viper.SetDefault("restserver.insecure-serving.port", a.cfg.RESTServer.InsecureServing.Port)
	viper.SetDefault("restserver.middlewares", a.cfg.RESTServer.Middlewares)
	viper.SetDefault("restserver.use-healthz", a.cfg.RESTServer.UseHealthz)

	viper.SetDefault("grpcserver.insecure-serving.port", a.cfg.GRPCServer.InsecureServing.Port)

	viper.SetTypeByDefaultValue(true)
}

func (a *app) bindConfigFromCliFlags() error {
	if err := viper.BindPFlags(a.cmd.Flags()); err != nil {
		msg := "failed to read config from cli flags"
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

func (a *app) bindConfig() error {
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

func (a *app) loadConfig() error {
	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal config"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *app) Run() {
	if err := a.cmd.Execute(); err != nil {
		zap.S().Fatal(err)
	}
}

func (a *app) runCmd(cmd *cobra.Command, args []string) error {
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
				Host: a.cfg.RESTServer.InsecureServing.Host,
				Port: a.cfg.RESTServer.InsecureServing.Port,
			},
			Mode:        a.cfg.RESTServer.Mode,
			Middlewares: a.cfg.RESTServer.Middlewares,
			UseHealthz:  a.cfg.RESTServer.UseHealthz,
		},
		serviceRouter,
	)
	go func() {
		if err := restServer.Start(); err != nil {
			msg := "error occured during running rest server"
			zap.S().Fatalw(msg, "err", err)
		}
	}()

	grpcServer := grpcserver.New(
		&grpcserver.Config{
			InsecureServing: &grpcserver.InsecureServingConfig{
				Host: a.cfg.GRPCServer.InsecureServing.Host,
				Port: a.cfg.GRPCServer.InsecureServing.Port,
			},
		},
		serviceRouter,
	)
	go func() {
		if err := grpcServer.Start(); err != nil {
			msg := "error occured during running grpc server"
			zap.S().Fatalw(msg, "err", err)
		}
	}()

	gracefulshutdown.Enable(func() error {
		waitDone := make(chan struct{}, 1)
		var shutdownErr error
		eg, ctx := errgroup.WithContext(context.Background())

		eg.Go(func() error {
			if err := restServer.Stop(); err != nil {
				shutdownErr = err
				return err
			}
			return nil
		})

		eg.Go(func() error {
			grpcServer.Stop()
			return nil
		})

		go func() {
			_ = eg.Wait()
			waitDone <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			return shutdownErr
		case <-waitDone:
			return nil
		}
	})

	select {}
}
