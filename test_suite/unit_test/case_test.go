package unit_test

import (
	"gopkg.in/yaml.v2"
	"os"
	"testing"
	"testrunner/common"
	"testrunner/test_report"
	"testrunner/test_suite"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

func (s *CaseSuite) SetUpTest(c *C) {
	_ = os.Setenv(common.PhaseEnv, string(common.Asserting))
	test_report.BuildReport()
	report := test_report.GetReport()
	report.AppendSuite("unit-test-suite")
}

func (s *CaseSuite) TestUnmarshalCaseFile(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_suite.Case{}

	// when
	err := yaml.Unmarshal(file, tc)

	// then
	c.Check(err, IsNil)
	c.Check(tc.Stages, HasLen, 3)
	c.Check(tc.Stages[0].Request.Body, Equals, "{\"sequence\": \"hello world\"}")
}

func (s *CaseSuite) TestCaseExecute(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_suite.Case{}

	// when
	_ = yaml.Unmarshal(file, tc)
	err := tc.Execute()

	// then
	c.Check(err, IsNil)
}

func (s *CaseSuite) TestCaseExecuteWithFilter(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_suite.Case{}

	// when
	_ = yaml.Unmarshal(file, tc)
	tc.AddFilter(func(stage test_suite.Stage) bool {
		return stage.Type == test_suite.API
	})
	err := tc.Execute()

	// then
	c.Check(err, IsNil)
}

func (s *CaseSuite) TestCaseExecuteWithReportGen(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_suite.Case{}

	// when
	_ = yaml.Unmarshal(file, tc)
	tc.AddFilter(func(stage test_suite.Stage) bool {
		return stage.Type == test_suite.API
	})
	err := tc.Execute()

	// then
	c.Check(err, IsNil)

	// when
	_, err = test_report.GetReport().Gen("gen")

	// then
	c.Check(err, IsNil)
}
