package config

import (
	"net"
	"strconv"
)

// GRPCServerConfig is the root struct for gRPC server related configs.
type GRPCServerConfig struct {
	Insecure *GRPCInsecureServingConfig `json:"insecure" mapstructure:"insecure"`
}

func newEmptyGRPCServerConfig() *GRPCServerConfig {
	return &GRPCServerConfig{
		Insecure: newEmptyGRPCInsecureServingConfig(),
	}
}

// GRPCInsecureServingConfig holds the configs for the insecure serving
// of gRPC server.
type GRPCInsecureServingConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

// Addr returns the full address(host:port).
func (c *GRPCInsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func newEmptyGRPCInsecureServingConfig() *GRPCInsecureServingConfig {
	return &GRPCInsecureServingConfig{}
}
