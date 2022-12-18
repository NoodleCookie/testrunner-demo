package test_case

type Executable interface {
	Execute() error
}

var _ Executable = (*Stage)(nil)
var _ Executable = (*Case)(nil)

type StageMold string

const (
	API = StageMold("api")
)
