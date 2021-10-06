package cliflags

import "github.com/spf13/pflag"

type grpcServerCliFlags struct {
	InsecureServing *GRPCInsecureServingCliFlags
	Mode            string

	flagSet *pflag.FlagSet
}

func newGrpcServerCliFlags() *grpcServerCliFlags {
	return &grpcServerCliFlags{
		InsecureServing: &GRPCInsecureServingCliFlags{},
	}
}

func (f *grpcServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("grpcserver", pflag.ExitOnError)

	fs.StringVar(&f.InsecureServing.Host, "grpcserver.insecure-serving.host", f.InsecureServing.Host, `GRPC server insecure serving host`)
	fs.IntVar(&f.InsecureServing.Port, "grpcserver.insecure-serving.port", f.InsecureServing.Port, `GRPC server insecure serving port`)

	f.flagSet = fs
	return f.flagSet
}

type GRPCInsecureServingCliFlags struct {
	Host string
	Port int
}
