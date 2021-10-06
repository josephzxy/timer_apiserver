package cliflags

import "github.com/spf13/pflag"

type CliFlags interface {
	GetAllFlagSets() []*pflag.FlagSet
}

type cliFlags struct {
	mysql      *mysqlCliFlags
	restserver *restServerCliFlags
	grpcserver *grpcServerCliFlags
}

func NewCliFlags() CliFlags {
	return &cliFlags{
		mysql:      newMysqlCliFlags(),
		restserver: newRestServerCliFlags(),
		grpcserver: newGrpcServerCliFlags(),
	}
}

func (f *cliFlags) GetAllFlagSets() []*pflag.FlagSet {
	return []*pflag.FlagSet{
		f.mysql.getFlagSet(),
		f.restserver.getFlagSet(),
		f.grpcserver.getFlagSet(),
	}
}
