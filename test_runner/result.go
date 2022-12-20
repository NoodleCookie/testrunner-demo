package test_runner

import (
	"net/http"
	"testrunner/test_case"
)

type Result struct {
	stage     test_case.Stage
	response  http.Response
	assertion Assertion
	pass      bool
}
