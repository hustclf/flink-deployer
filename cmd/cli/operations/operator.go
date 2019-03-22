package operations

import (
	"flink-deployer/cmd/cli/flink"
	"github.com/spf13/afero"
)

// Operator is an interface which contains all the functionality
// that the deployer exposes
type Operator interface {
	Deploy(d Deploy) error
	Update(u UpdateJob) error
	Terminate(t TerminateJob) error
	RetrieveJobs() ([]flink.Job, error)
}

// RealOperator is the Operator used in the production code
type RealOperator struct {
	Filesystem   afero.Fs
	FlinkRestAPI flink.FlinkRestAPI
}
