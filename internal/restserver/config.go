package restserver

import (
	"net"
	"strconv"
)

// Config holds configs for REST server.
type Config struct {
	InsecureServing InsecureServingConfig
	Mode            string
	Middlewares     []string
	UseHealthz      bool
}

// InsecureServingConfig holds insecure serving configs for REST server.
type InsecureServingConfig struct {
	Host string
	Port int
}

// Addr returns the full address(host:port)
func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
