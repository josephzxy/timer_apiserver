package config

import "time"

type MySQLConfig struct {
	User            string        `json:"user" mapstructure:"user"`
	Pwd             string        `json:"pwd" mapstructure:"pwd"`
	Host            string        `json:"host" mapstructure:"host"`
	Port            int           `json:"port" mapstructure:"port"`
	Database        string        `json:"database" mapstructure:"database"`
	Charset         string        `json:"charset" mapstructure:"charset"`
	ParseTime       bool          `json:"parse-time" mapstructure:"parse-time"`
	Loc             string        `json:"loc" mapstructure:"loc"`
	MaxIdleConns    int           `json:"max-idle-conns" mapstructure:"max-idle-conns"`
	MaxOpenConns    int           `json:"max-open-conns" mapstructure:"max-open-conns"`
	MaxConnLifetime time.Duration `json:"max-conn-lifetime" mapstructure:"max-conn-lifetime"`
	LogLevel        int           `json:"log-level" mapstructure:"log-level"`
}

func newEmptyMySQLConfig() *MySQLConfig {
	return &MySQLConfig{}
}
