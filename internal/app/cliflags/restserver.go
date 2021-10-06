package cliflags

import "github.com/spf13/pflag"

type restServerCliFlags struct {
	Insecure    *RESTInsecureServingCliFlags
	Mode        string
	Middlewares []string
	UseHealthz  bool

	flagSet *pflag.FlagSet
}

func newRestServerCliFlags() *restServerCliFlags {
	return &restServerCliFlags{
		Insecure: &RESTInsecureServingCliFlags{},
	}
}

func (f *restServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("restserver", pflag.ExitOnError)

	fs.StringVar(&f.Insecure.Host, "rest.insecure.host", f.Insecure.Host, `REST server insecure serving host`)
	fs.IntVar(&f.Insecure.Port, "rest.insecure.port", f.Insecure.Port, `REST server insecure serving port`)
	fs.StringVar(&f.Mode, "rest.mode", f.Mode, `REST server mode(e.g. debug, release)`)
	fs.StringArrayVar(&f.Middlewares, "rest.middlewares", f.Middlewares, `REST server middlewares`)
	fs.BoolVar(&f.UseHealthz, "rest.use-healthz", f.UseHealthz, `Whether or not to enable health checking endpoint /healthz for REST server`)

	f.flagSet = fs
	return f.flagSet
}

type RESTInsecureServingCliFlags struct {
	Host string
	Port int
}
