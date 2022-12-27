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

type Actual struct {
	Status  string
	Headers map[string]string
	Body    string
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
		request, err := http.NewRequest(s.Request.Method, s.Request.Url, nil)
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}
		statusCode := res.StatusCode
		body := res.Body
		if err != nil {
			panic(err)
		}

		defer body.Close()

		if s.Assertion != nil {
			assertor := test_assertion.GenericAssertor{}
			assert, err := assertor.Assert(statusCode, test_assertion.Equals, s.Assertion.Status)
			if err != nil || !assert {
				s.apiReport(assert, err.Error())
			} else {
				s.apiReport(assert, map[string]any{"request": s.Request, "assertion": s.Assertion})
			}
		}

		return nil
	}

	return nil
}

func (s *Stage) apiReport(pass bool, detail any) {
	if common.CurrentPhase() == common.Asserting {
		test_report.GetReport().AppendStage(s.Name, pass, detail)
	}
}
