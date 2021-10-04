package config

type Config struct {
	MySQL      *MySQLConfig      `mapstructure:"mysql"`
	RESTServer *RESTServerConfig `mapstructure:"restserver"`
}
