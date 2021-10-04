package config

type Config struct {
	MySQL      *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	RESTServer *RESTServerConfig `json:"restserver" mapstructure:"restserver"`
}
