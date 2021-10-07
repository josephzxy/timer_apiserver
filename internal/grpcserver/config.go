package grpcserver

import (
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// Config holds configs for gRPC server
type Config struct {
	InsecureServing *InsecureServingConfig
}

// InsecureServingConfig holds insecure serving configs for gRPC server
type InsecureServingConfig struct {
	Host    string
	Port    int
	Options []grpc.ServerOption
}

// Addr returns the full address(host:port)
func (c *InsecureServingConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
