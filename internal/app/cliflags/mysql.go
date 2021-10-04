package cliflags

import (
	"time"

	"github.com/spf13/pflag"
)

type MySQLCliFlags struct {
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

	flagSet *pflag.FlagSet
}

func newMySQLCliFlags() *MySQLCliFlags {
	return &MySQLCliFlags{}
}

func (f *MySQLCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("mysql", pflag.ExitOnError)

	fs.StringVar(&f.User, "mysql.user", f.User, `MySQL service user`)
	fs.StringVar(&f.Pwd, "mysql.pwd", f.Pwd, `MySQL service password`)
	fs.StringVar(&f.Host, "mysql.host", f.Host, `MySQL service host address`)
	fs.IntVar(&f.Port, "mysql.port", f.Port, `MySQL service port`)
	fs.StringVar(&f.Database, "mysql.database", f.Database, `MySQL service database`)
	fs.StringVar(&f.Charset, "mysql.charset", f.Charset, `MySQL service charset`)
	fs.BoolVar(&f.ParseTime, "mysql.parse-time", f.ParseTime, `Whether or not to parse time in the query`)
	fs.StringVar(&f.Loc, "mysql.loc", f.Loc, `The timezone used in the query to MySQL server`)
	fs.IntVar(&f.MaxIdleConns, "mysql.max-idle-conns", f.MaxIdleConns, `The max number of idle connections allowed to the MySQL server`)
	fs.IntVar(&f.MaxOpenConns, "mysql.max-open-conns", f.MaxOpenConns, `The max number of open connections allowed to the MySQL server`)
	fs.DurationVar(&f.MaxConnLifetime, "mysql.max-conn-lifetime", f.MaxConnLifetime, `The max lifetime of a connection to the MySQL server`)
	fs.IntVar(&f.LogLevel, "mysql.log-level", f.LogLevel, `Gorm log level`)

	f.flagSet = fs
	return f.flagSet
}
