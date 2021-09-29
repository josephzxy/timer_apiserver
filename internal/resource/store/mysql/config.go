package mysql

import "time"

type Config struct {
	User            string
	Pwd             string
	Host            string
	Port            int
	Database        string
	Charset         string
	ParseTime       bool
	Loc             string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxConnLifetime time.Duration
	LogLevel        int
}
