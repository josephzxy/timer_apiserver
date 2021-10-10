package cliflags

import (
	"time"

	"github.com/spf13/pflag"
)

type mysqlCliFlags struct {
	user            string
	pwd             string
	host            string
	port            int
	database        string
	charset         string
	parseTime       bool
	loc             string
	maxIdleConns    int
	maxOpenConns    int
	maxConnLifetime time.Duration
	logLevel        int

	flagSet *pflag.FlagSet
}

func newMysqlCliFlags() *mysqlCliFlags {
	return &mysqlCliFlags{}
}

func (f *mysqlCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("mysql", pflag.ExitOnError)

	fs.String("mysql.user", f.user, `MySQL service user`)
	fs.String("mysql.pwd", f.pwd, `MySQL service password`)
	fs.String("mysql.host", f.host, `MySQL service host address`)
	fs.Int("mysql.port", f.port, `MySQL service port`)
	fs.String("mysql.database", f.database, `MySQL service database`)
	fs.String("mysql.charset", f.charset, `MySQL service charset`)
	fs.Bool("mysql.parse-time", f.parseTime, `Whether or not to parse time in the query`)
	fs.String("mysql.loc", f.loc, `The timezone used in the query to MySQL server`)
	fs.Int("mysql.max-idle-conns", f.maxIdleConns, `The max number of idle connections allowed to the MySQL server`)
	fs.Int("mysql.max-open-conns", f.maxOpenConns, `The max number of open connections allowed to the MySQL server`)
	fs.Duration("mysql.max-conn-lifetime", f.maxConnLifetime, `The max lifetime of a connection to the MySQL server`)
	fs.Int("mysql.log-level", f.logLevel, `Gorm log level`)

	f.flagSet = fs
	return f.flagSet
}
