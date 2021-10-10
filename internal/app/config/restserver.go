package config

import (
	"net"
	"strconv"
)

// RESTServerConfig is the root struct for REST server related configs.
type RESTServerConfig struct {
	Insecure    *RESTInsecureServingConfig `json:"insecure" mapstructure:"insecure"`
	Mode        string                     `json:"mode" mapstructure:"mode"`
	Middlewares []string                   `json:"middlewares" mapstructure:"middlewares"`
	UseHealthz  bool                       `json:"use-healthz" mapstructure:"use-healthz"`
}

func newEmptyRESTServerConfig() *RESTServerConfig {
	return &RESTServerConfig{
		Insecure: newEmptyRESTInsecureServingConfig(),
	}
}

// RESTInsecureServingConfig holds the configs for the insecure serving
// of REST server.
type RESTInsecureServingConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

// Addr returns the full address(host:port).
func (c *RESTInsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func newEmptyRESTInsecureServingConfig() *RESTInsecureServingConfig {
	return &RESTInsecureServingConfig{}
}
