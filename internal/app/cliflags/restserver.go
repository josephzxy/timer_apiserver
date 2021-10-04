package cliflags

import "github.com/spf13/pflag"

type RESTServerCliFlags struct {
	InsecureServing *InsecureServingCliFlags
	Mode            string

	flagSet *pflag.FlagSet
}

func newRESTServerCliFlags() *RESTServerCliFlags {
	return &RESTServerCliFlags{
		InsecureServing: &InsecureServingCliFlags{},
	}
}

func (f *RESTServerCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("restserver", pflag.ExitOnError)

	fs.StringVar(&f.InsecureServing.Host, "restserver.insecure-serving.host", f.InsecureServing.Host, `REST server insecure serving host`)
	fs.IntVar(&f.InsecureServing.Port, "restserver.insecure-serving.port", f.InsecureServing.Port, `REST server insecure serving port`)
	fs.StringVar(&f.Mode, "restserver.mode", f.Mode, `REST server mode(e.g. debug, release)`)

	f.flagSet = fs
	return f.flagSet
}

type InsecureServingCliFlags struct {
	Host string
	Port int
}
