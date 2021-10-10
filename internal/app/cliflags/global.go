package cliflags

import "github.com/spf13/pflag"

type globalCliFlags struct {
	config string

	flagSet *pflag.FlagSet
}

func newGlobalCliFlags() *globalCliFlags {
	return &globalCliFlags{}
}

func (f *globalCliFlags) getFlagSet() *pflag.FlagSet {
	if f.flagSet != nil {
		return f.flagSet
	}
	fs := pflag.NewFlagSet("global", pflag.ExitOnError)

	fs.String("config", f.config, `The path to config file. Reading from config file will be skipped if path not set`)

	f.flagSet = fs
	return f.flagSet
}
