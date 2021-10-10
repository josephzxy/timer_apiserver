package cliflags

import "github.com/spf13/pflag"

type grpcServerCliFlags struct {
	insecure *grpcInsecureServingCliFlags
	mode     string

	flagSet *pflag.FlagSet
}

func newGrpcServerCliFlags() *grpcServerCliFlags {
	return &grpcServerCliFlags{
		insecure: &grpcInsecureServingCliFlags{},
	}
}

func (f *grpcServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("grpcserver", pflag.ExitOnError)

	fs.String("grpc.insecure.host", f.insecure.host, `GRPC server insecure serving host`)
	fs.Int("grpc.insecure.port", f.insecure.port, `GRPC server insecure serving port`)
	fs.String("grpc.mode", f.mode, `GRPC server mode`)

	f.flagSet = fs
	return f.flagSet
}

type grpcInsecureServingCliFlags struct {
	host string
	port int
}
