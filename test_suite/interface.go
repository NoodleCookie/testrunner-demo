package test_suite

type Executable interface {
	Execute() error
}

var _ Executable = (*Stage)(nil)
var _ Executable = (*Case)(nil)
var _ Executable = (*Suite)(nil)

type StageType string

const (
	API = StageType("api")
)
