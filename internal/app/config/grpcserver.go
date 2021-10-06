package config

import (
	"net"
	"strconv"
)

type GRPCServerConfig struct {
	Insecure *GRPCInsecureServingConfig `json:"insecure" mapstructure:"insecure"`
}

func newEmptyGRPCServerConfig() *GRPCServerConfig {
	return &GRPCServerConfig{
		Insecure: newEmptyGRPCInsecureServingConfig(),
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
