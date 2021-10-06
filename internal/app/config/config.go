package config

type Config struct {
	MySQL      *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	RESTServer *RESTServerConfig `json:"restserver" mapstructure:"restserver"`
	GRPCServer *GRPCServerConfig `json:"grpcserver" mapstructure:"grpcserver"`

	Config string `json:"config" mapstructure:"config"`
}

func NewEmptyConfig() *Config {
	return &Config{
		MySQL:      newEmptyMySQLConfig(),
		RESTServer: newEmptyRESTServerConfig(),
		GRPCServer: newEmptyGRPCServerConfig(),
	}
}
