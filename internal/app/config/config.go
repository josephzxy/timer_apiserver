package config

var cfgFile = "config/config.yml"

type Config struct {
	MySQL      *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	RESTServer *RESTServerConfig `json:"restserver" mapstructure:"restserver"`
	GRPCServer *GRPCServerConfig `json:"grpcserver" mapstructure:"grpcserver"`
}

func CfgFile() string {
	return cfgFile
}

func NewEmptyConfig() *Config {
	return &Config{
		MySQL:      newEmptyMySQLConfig(),
		RESTServer: newEmptyRESTServerConfig(),
		GRPCServer: newEmptyGRPCServerConfig(),
	}
}
