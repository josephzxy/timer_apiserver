package config

type Config struct {
	MySQL *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	REST  *RESTServerConfig `json:"rest" mapstructure:"rest"`
	GRPC  *GRPCServerConfig `json:"grpc" mapstructure:"grpc"`

	Config string `json:"config" mapstructure:"config"`
}

func NewEmptyConfig() *Config {
	return &Config{
		MySQL: newEmptyMySQLConfig(),
		REST:  newEmptyRESTServerConfig(),
		GRPC:  newEmptyGRPCServerConfig(),
	}
}
