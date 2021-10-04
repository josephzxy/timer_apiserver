package config

import (
	"net"
	"strconv"
)

type RESTServerConfig struct {
	InsecureServing *InsecureServingConfig `json:"insecure-serving" mapstructure:"insecure-serving"`
	Mode            string                 `json:"mode" mapstructure:"mode"`
}

func newEmptyRESTServerConfig() *RESTServerConfig {
	return &RESTServerConfig{
		InsecureServing: newEmptyInsecureServingConfig(),
	}
}

type InsecureServingConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func newEmptyInsecureServingConfig() *InsecureServingConfig {
	return &InsecureServingConfig{}
}
