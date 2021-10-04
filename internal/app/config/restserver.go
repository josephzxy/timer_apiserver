package config

import (
	"net"
	"strconv"
)

type RESTServerConfig struct {
	InsecureServing *InsecureServingConfig `mapstructure:"insecure-serving"`
	Mode            string                 `mapstructure:"mode"`
}

type InsecureServingConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
