package config

import (
	"net"
	"strconv"
)

type RESTServerConfig struct {
	InsecureServing *RESTInsecureServingConfig `json:"insecure-serving" mapstructure:"insecure-serving"`
	Mode            string                     `json:"mode" mapstructure:"mode"`
	Middlewares     []string                   `json:"middlewares" mapstructure:"middlewares"`
}

func newEmptyRESTServerConfig() *RESTServerConfig {
	return &RESTServerConfig{
		InsecureServing: newEmptyRESTInsecureServingConfig(),
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
