// Package cliflags manages all command line flags for the app
package cliflags

import "github.com/spf13/pflag"

// Cliflags defines the behavior of the app-level
// command line flags manager
type CliFlags interface {
	GetAllFlagSets() []*pflag.FlagSet
}

type cliFlags struct {
	mysql      *mysqlCliFlags
	restserver *restServerCliFlags
	grpcserver *grpcServerCliFlags
	global     *globalCliFlags
}

// NewCliFlags returns a value of an implementation
// of interface CliFlags
func NewCliFlags() CliFlags {
	return &cliFlags{
		mysql:      newMysqlCliFlags(),
		restserver: newRestServerCliFlags(),
		grpcserver: newGrpcServerCliFlags(),
		global:     newGlobalCliFlags(),
	}
}

// GetAllFlagSets returns all flag sets defined in
// the implementation of interface CliFlags
func (f *cliFlags) GetAllFlagSets() []*pflag.FlagSet {
	return []*pflag.FlagSet{
		f.mysql.getFlagSet(),
		f.restserver.getFlagSet(),
		f.grpcserver.getFlagSet(),
		f.global.getFlagSet(),
	}
}
