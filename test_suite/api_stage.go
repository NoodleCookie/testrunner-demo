package test_suite

import (
	"net/http"
	"testrunner/common"
	"testrunner/test_assertion"
	"testrunner/test_report"
)

type Request struct {
	Url     string            `yaml:"url,omitempty"`
	Method  string            `yaml:"method,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty"`
	Body    string            `yaml:"body,omitempty"`
}

type Assertion struct {
	Status  int               `yaml:"status,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    struct {
		Equals string `yaml:"equals,omitempty"`
	} `yaml:"body,omitempty"`
}

func (s *Stage) executeApi() error {
	if common.CurrentPhase() == common.Asserting {
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
		if err != nil {
			panic(err)
		}

		defer body.Close()

		assertor := test_assertion.GenericAssertor{}

		assert, err := assertor.Assert(statusCode, test_assertion.Equals, s.Assertion.Status)
		if err != nil || !assert {
			s.apiReport(assert, err.Error())
		} else {
			s.apiReport(assert, map[string]interface{}{"request": s.Request, "assertion": s.Assertion})
		}

		return nil
	}

	return nil
}

func (s *Stage) apiReport(pass bool, detail interface{}) {
	if common.CurrentPhase() == common.Asserting {
		test_report.GetReport().AppendStage(s.Name, pass, detail)
	}
}
