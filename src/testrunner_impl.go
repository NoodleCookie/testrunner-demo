package src

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

type SimpleTestrunner struct {
}

func (s SimpleTestrunner) Resolve(file string) (Testcase, error) {
	testcaseBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("resolve %s", err.Error())
	}
	testcase := SimpleTestcase{}
	if err := yaml.Unmarshal(testcaseBytes, &testcase); err != nil {
		return nil, fmt.Errorf("resolve %s", err.Error())
	}
	return testcase, nil
}

func (s SimpleTestrunner) Send(testcase Testcase) (Testcase, error) {
	client := http.DefaultClient
	request := testcase.Request()
	response, err := client.Do(&request)
	if err != nil {
		return nil, err
	}
	return SimpleTestcase{
		Res: *response,
		Req: testcase.Request(),
	}, nil
}

func (s SimpleTestrunner) Verify(real http.Response, expect Testcase) (bool, string, error) {
	realCode := real.Status
	expectCode := expect.Response().Status
	if realCode == expectCode {
		return true, "succeed", nil
	}
	return false, "failed", errors.New(fmt.Sprintf("real status code %s not match expect code %s", realCode, expectCode))
}
