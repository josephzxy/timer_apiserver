package cliflags

import "github.com/spf13/pflag"

type CliFlags struct {
	mysql      *MySQLCliFlags
	restserver *RESTServerCliFlags
}

func NewCliFlags() *CliFlags {
	return &CliFlags{
		mysql:      newMySQLCliFlags(),
		restserver: newRESTServerCliFlags(),
	}
}

func (f *CliFlags) GetAllFlagSets() []*pflag.FlagSet {
	return []*pflag.FlagSet{
		f.mysql.getFlagSet(),
		f.restserver.getFlagSet(),
	}
}
