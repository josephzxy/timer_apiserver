package config

import (
	"net"
	"strconv"
)

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

type RESTInsecureServingConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func (c *RESTInsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func newEmptyRESTInsecureServingConfig() *RESTInsecureServingConfig {
	return &RESTInsecureServingConfig{}
}
