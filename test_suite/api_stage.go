package test_suite

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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
	Status  *int
	Headers map[string]string
	Body    *string
}

type Assertion struct {
	Status  int               `yaml:"status,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    struct {
		Equals string `yaml:"equals,omitempty"`
	} `yaml:"body,omitempty"`
}

func (s *Stage) executeApi() error {
	request, err := http.NewRequest(s.getRenderRequest().Method, s.getRenderRequest().Url, nil)
	if err != nil {
		panic(err)
	}
	request.Header = recoverHeaders(s.getRenderRequest().Headers)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	statusCode := resp.StatusCode
	body := resp.Body
	if err != nil {
		panic(err)
	}

	defer body.Close()

	s.Actual = &Actual{
		Status:  &resp.StatusCode,
		Headers: transferHeaders(resp.Header),
		Body:    transferBody(resp.Body),
	}

	if common.CurrentPhase() == common.Asserting {
		if s.Assertion != nil {
			assertor := test_assertion.GenericAssertor{}
			assert, err := assertor.Assert(statusCode, test_assertion.Equals, s.Assertion.Status)
			if err != nil || !assert {
				s.apiReport(assert, err.Error())
			} else {
				s.apiReport(assert, map[string]any{"request": s.getRenderRequest(), "assertion": s.Assertion})
			}
		}
	}

	return nil
}

func transferBody(body io.ReadCloser) *string {
	if body == nil {
		return nil
	}
	content, err := io.ReadAll(body)
	if err != nil {
		msg := fmt.Sprintf("transfer body error: %s", err.Error())
		return &msg
	}
	result := string(content)
	return &result
}

func transferHeaders(header http.Header) map[string]string {
	if header == nil {
		return nil
	}
	result := make(map[string]string, 0)
	for key, values := range header {
		result[key] = strings.Join(values, ", ")
	}
	return result
}

func recoverHeaders(header map[string]string) http.Header {
	if header == nil {
		return nil
	}
	result := make(map[string][]string, 0)
	for key, values := range header {
		result[key] = strings.Split(values, ", ")
	}
	return result
}

func (s *Stage) apiReport(pass bool, detail any) {
	if common.CurrentPhase() == common.Asserting {
		test_report.GetReport().AppendStage(s.Name, pass, detail)
	}
}
