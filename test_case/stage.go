package test_case

import (
	"fmt"
	"testrunner/test_runner"
)

type Stage struct {
	Mold    StageMold          `yaml:"type,omitempty"`
	Request Request            `yaml:"request,omitempty"`
	Expect  test_runner.Expect `yaml:"expect,omitempty"`
}

type Request struct {
	Url     string            `yaml:"url,omitempty"`
	Method  string            `yaml:"method,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty"`
	Body    string            `yaml:"body,omitempty"`
}

func (s *Stage) Execute() error {
	// todo: implement
	fmt.Println("execute request: ", s.Request.Url)
	return nil
}
