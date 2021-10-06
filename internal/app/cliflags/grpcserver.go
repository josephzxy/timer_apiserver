package cliflags

import "github.com/spf13/pflag"

type grpcServerCliFlags struct {
	Insecure *GRPCInsecureServingCliFlags
	Mode     string

	flagSet *pflag.FlagSet
}

func newGrpcServerCliFlags() *grpcServerCliFlags {
	return &grpcServerCliFlags{
		Insecure: &GRPCInsecureServingCliFlags{},
	}
}

func (f *grpcServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("grpcserver", pflag.ExitOnError)

	fs.StringVar(&f.Insecure.Host, "grpc.insecure.host", f.Insecure.Host, `GRPC server insecure serving host`)
	fs.IntVar(&f.Insecure.Port, "grpc.insecure.port", f.Insecure.Port, `GRPC server insecure serving port`)

	f.flagSet = fs
	return f.flagSet
}

type GRPCInsecureServingCliFlags struct {
	Host string
	Port int
}
