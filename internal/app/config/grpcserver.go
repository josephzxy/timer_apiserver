package config

import (
	"net"
	"strconv"
)

type GRPCServerConfig struct {
	InsecureServing *GRPCInsecureServingConfig `json:"insecure-serving" mapstructure:"insecure-serving"`
}

func newEmptyGRPCServerConfig() *GRPCServerConfig {
	return &GRPCServerConfig{
		InsecureServing: newEmptyGRPCInsecureServingConfig(),
	}
}

type GRPCInsecureServingConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func (c *GRPCInsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func newEmptyGRPCInsecureServingConfig() *GRPCInsecureServingConfig {
	return &GRPCInsecureServingConfig{}
}
