// Package config manages all app-level configs
package config

// Config is the root struct that holds all app-level configs.
// One common use case is to create an empty value of Config
// and unmarshal values into it from sources like viper so that
// app-level configs don't vary throughout the whole lifecycle
// of the app.
type Config struct {
	MySQL *MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	REST  *RESTServerConfig `json:"rest" mapstructure:"rest"`
	GRPC  *GRPCServerConfig `json:"grpc" mapstructure:"grpc"`

	Config string `json:"config" mapstructure:"config"`
}

// NewEmptyConfig returns an empty value of Config
func NewEmptyConfig() *Config {
	return &Config{
		MySQL: newEmptyMySQLConfig(),
		REST:  newEmptyRESTServerConfig(),
		GRPC:  newEmptyGRPCServerConfig(),
	}
}
