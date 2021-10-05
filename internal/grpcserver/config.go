package grpcserver

import (
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Config struct {
	InsecureServing *InsecureServingConfig
}

type InsecureServingConfig struct {
	Host    string
	Port    int
	Options []grpc.ServerOption
}

func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
