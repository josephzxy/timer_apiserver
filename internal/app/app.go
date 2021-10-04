package app

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/josephzxy/timer_apiserver/internal/app/cliflags"
	"github.com/josephzxy/timer_apiserver/internal/app/config"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql"
	"github.com/josephzxy/timer_apiserver/internal/restserver"
)

type App struct {
	basename string
	cfg      *config.Config
	cliflags *cliflags.CliFlags
}

func NewApp(basename string) *App {
	a := &App{
		basename: basename,
		cfg:      config.NewEmptyConfig(),
		cliflags: cliflags.NewCliFlags(),
	}

	a.installCliFlags()
	if err := a.loadConfig(); err != nil {
		zap.S().Fatalw("failed to load config for app", "err", err)
	}
	return a
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

func (a *App) loadConfigFromCliFlags() error {
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		msg := "failed to read config from cli flags"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) loadConfigFromEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ToUpper(a.basename))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func (a *App) loadConfigFromFile() error {
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

func (a *App) loadConfig() error {
	a.ensureViperValueType()
	if err := a.loadConfigFromCliFlags(); err != nil {
		return err
	}
	a.loadConfigFromEnv()
	if err := a.loadConfigFromFile(); err != nil {
		return err
	}
	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal config"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) installCliFlags() {
	for _, fs := range a.cliflags.GetAllFlagSets() {
		pflag.CommandLine.AddFlagSet(fs)
	}
	pflag.Parse()
}

func (a *App) Run() {
	fmt.Println("Running...")

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
		zap.S().Fatalw("failed to get mysql store router", "err", err)
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

	if err := restServer.Start(); err != nil {
		zap.S().Fatalw("error occured during running rest server", "err", err)
	}
}
