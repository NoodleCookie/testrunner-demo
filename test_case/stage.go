package test_case

import (
	"io"
	"net/http"
)

type Stage struct {
	Mold    StageMold `yaml:"type,omitempty"`
	Request Request   `yaml:"request,omitempty"`
	Expect  Response  `yaml:"expect,omitempty"`
}

type Request struct {
	Url     string            `yaml:"url,omitempty"`
	Method  string            `yaml:"method,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty"`
	Body    string            `yaml:"body,omitempty"`
}

func (s *Stage) Execute() error {
	//send http request base on stage request
	request, err := http.NewRequest(s.Request.Method, s.Request.Url, nil)
	if err != nil {
		panic(err)
	}
	//get http response
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	//get status code from response
	statusCode := res.StatusCode
	//get body from response
	body := res.Body
	resByte, err := io.ReadAll(body)
	resString := string(resByte)
	//compare with expect and generate pass
	pass := (s.Expect.Res == resString) && (s.Expect.Status == statusCode)
	//generate compare
	compare := Compare{expect: s.Expect, actual: Response{statusCode, resString}}
	//return result
	result := Result{s.Request, Response{statusCode, resString}, compare, pass}
	results := append(report.Results, result)
	report.Results = results[:]
	return nil
}
