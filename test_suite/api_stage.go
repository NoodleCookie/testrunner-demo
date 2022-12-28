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
	Url     string            `yaml:"url,omitempty" json:"url,omitempty"`
	Method  string            `yaml:"method,omitempty" json:"method,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty" json:"params,omitempty"`
	Body    string            `yaml:"body,omitempty" json:"body,omitempty"`
}

type Actual struct {
	Status  *int              `yaml:"status,omitempty" json:"status,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	Body    *string           `yaml:"body,omitempty" json:"body,omitempty"`
}

type Assertion struct {
	Status  int               `yaml:"status,omitempty" json:"status,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	Body    string            `yaml:"body,omitempty" json:"body,omitempty"`
}

func (a *Actual) Success(status int, headers map[string]string, body *string) *Actual {
	a.Status = &status
	a.Headers = headers
	a.Body = body
	return a
}

func (s *Stage) executeApi() error {

	s.Actual = &Actual{}

	request, err := http.NewRequest(s.getRenderRequest().Method, s.getRenderRequest().Url, nil)
	if err != nil {
		s.report(false, err.Error())
		return nil
	}
	request.Header = recoverHeaders(s.getRenderRequest().Headers)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		s.report(false, err.Error())
		return nil
	}
	statusCode := resp.StatusCode
	body := resp.Body
	if err != nil {
		s.report(false, err.Error())
		return nil
	}

	defer body.Close()

	s.Actual.Success(resp.StatusCode, transferHeaders(resp.Header), transferBody(resp.Body))

	if common.CurrentPhase() == common.Asserting && s.Assertion != nil {
		assertor := test_assertion.GenericAssertor{}
		assert, err := assertor.Assert(statusCode, test_assertion.Equals, s.Assertion.Status)
		if err != nil || !assert {
			s.report(assert, err.Error())
		} else {
			s.report(assert, map[string]any{"request": s.getRenderRequest(), "assertion": s.Assertion})
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

func (s *Stage) report(pass bool, detail any) {
	test_report.GetReport().AppendStage(s.Name, pass, detail)
}
