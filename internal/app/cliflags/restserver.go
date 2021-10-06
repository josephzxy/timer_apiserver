package cliflags

import "github.com/spf13/pflag"

type restServerCliFlags struct {
	InsecureServing *RESTInsecureServingCliFlags
	Mode            string
	Middlewares     []string
	UseHealthz      bool

	flagSet *pflag.FlagSet
}

func newRestServerCliFlags() *restServerCliFlags {
	return &restServerCliFlags{
		InsecureServing: &RESTInsecureServingCliFlags{},
	}
}

func (f *restServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("restserver", pflag.ExitOnError)

	fs.StringVar(&f.InsecureServing.Host, "restserver.insecure-serving.host", f.InsecureServing.Host, `REST server insecure serving host`)
	fs.IntVar(&f.InsecureServing.Port, "restserver.insecure-serving.port", f.InsecureServing.Port, `REST server insecure serving port`)
	fs.StringVar(&f.Mode, "restserver.mode", f.Mode, `REST server mode(e.g. debug, release)`)
	fs.StringArrayVar(&f.Middlewares, "restserver.middlewares", f.Middlewares, `REST server middlewares`)
	fs.BoolVar(&f.UseHealthz, "restserver.use-healthz", f.UseHealthz, `Whether or not to enable health checking endpoint /healthz for REST server`)

	f.flagSet = fs
	return f.flagSet
}

type RESTInsecureServingCliFlags struct {
	Host string
	Port int
}
