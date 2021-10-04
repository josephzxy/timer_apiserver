package config

import "time"

type MySQLConfig struct {
	User            string        `mapstructure:"user"`
	Pwd             string        `mapstructure:"pwd"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Database        string        `mapstructure:"database"`
	Charset         string        `mapstructure:"charset"`
	ParseTime       bool          `mapstructure:"parse-time"`
	Loc             string        `mapstructure:"loc"`
	MaxIdleConns    int           `mapstructure:"max-idle-conns"`
	MaxOpenConns    int           `mapstructure:"max-open-conns"`
	MaxConnLifetime time.Duration `mapstructure:"max-conn-lifetime"`
	LogLevel        int           `mapstructure:"log-level"`
}
