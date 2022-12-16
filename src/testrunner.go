package src

import "net/http"

type Testrunner interface {
	Resolve(file string) (Testcase, error)

	Send(testcase Testcase) (Testcase, error)

	Verify(real http.Response, expect Testcase) (bool, string, error)
}

type Testcase interface {
	Request() http.Request
	Response() http.Response
}

var _ Testrunner = (*SimpleTestrunner)(nil)
var _ Testcase = (*SimpleTestcase)(nil)
