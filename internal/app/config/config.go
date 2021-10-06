package config

type Config struct {
	MySQL      *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	REST       *RESTServerConfig `json:"rest" mapstructure:"rest"`
	GRPCServer *GRPCServerConfig `json:"grpcserver" mapstructure:"grpcserver"`

	Config string `json:"config" mapstructure:"config"`
}

func NewEmptyConfig() *Config {
	return &Config{
		MySQL:      newEmptyMySQLConfig(),
		REST:       newEmptyRESTServerConfig(),
		GRPCServer: newEmptyGRPCServerConfig(),
	}
}
