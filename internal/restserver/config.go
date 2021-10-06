package restserver

import (
	"net"
	"strconv"
)

type Config struct {
	InsecureServing InsecureServingConfig
	Mode            string
	Middlewares     []string
}

type InsecureServingConfig struct {
	Host string
	Port int
}

func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
