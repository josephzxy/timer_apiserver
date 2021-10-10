package cliflags

import "github.com/spf13/pflag"

type restServerCliFlags struct {
	insecure    *restInsecureServingCliFlags
	mode        string
	middlewares []string
	useHealthz  bool

	flagSet *pflag.FlagSet
}

func newRestServerCliFlags() *restServerCliFlags {
	return &restServerCliFlags{
		insecure: &restInsecureServingCliFlags{},
	}
}

func (f *restServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("restserver", pflag.ExitOnError)

	fs.String("rest.insecure.host", f.insecure.host, `REST server insecure serving host`)
	fs.Int("rest.insecure.port", f.insecure.port, `REST server insecure serving port`)
	fs.String("rest.mode", f.mode, `REST server mode(e.g. debug, release)`)
	fs.StringArray("rest.middlewares", f.middlewares, `REST server middlewares`)
	fs.Bool("rest.use-healthz", f.useHealthz, `Whether or not to enable health checking endpoint /healthz for REST server`)

	f.flagSet = fs
	return f.flagSet
}

type restInsecureServingCliFlags struct {
	host string
	port int
}
