package unit_test

import (
	"encoding/json"
	"os"
	"testing"
	"testrunner/test_report"
	"testrunner/test_runner"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

//func (s *CaseSuite) TestRunnerCorrectCheck(c *C) {
//	// given
//	testrunner := test_runner.Testrunner{}
//
//	// when
//	description, err := testrunner.CheckDescription("./data/testsuite_correct_import")
//
//	// then
//	c.Check(err, IsNil)
//	c.Check(description.Import, HasLen, 1)
//	c.Check(description.Import[0], Equals, "hello-test")
//}
//
//func (s *CaseSuite) TestRunnerErrorCheck(c *C) {
//	// given
//	testrunner := test_runner.Testrunner{}
//
//	// when
//	_, err := testrunner.CheckDescription("./data/testsuite_nil_description")
//
//	// then
//	c.Check(err, NotNil)
//	c.Assert(err.Error(), Equals, "you must import your testcase into description")
//}

func (s *CaseSuite) TestRunnerRun(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_correct_import")

	// then
	c.Check(err, IsNil)
}

func (s *CaseSuite) TestRunnerExec(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_correct_baidu_request")

	// then
	c.Check(err, IsNil)

	// when
	_, err = test_report.GetReport().Gen("./gen")

	// then
	c.Check(err, IsNil)
}

func (s *CaseSuite) TestRunnerReportWithCorrectRequest(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_correct_baidu_request")

	// then
	c.Check(err, IsNil)

	// when
	_, err = test_report.GetReport().Gen("./gen")

	// then
	file, _ := os.ReadFile("./gen/testrunner-report.json")
	report := test_report.Report{}
	_ = json.Unmarshal([]byte(file), &report)
	c.Check(report, NotNil)
	c.Check(report.Pass, Equals, true)
	c.Check(report.Suites, HasLen, 1)
	c.Check(report.Suites[0].Pass, Equals, true)
	c.Check(report.Suites[0].Cases, HasLen, 2)
	c.Check(report.Suites[0].Cases[0].RunTime, NotNil)
	c.Check(report.Suites[0].Cases[0].Pass, Equals, true)
	c.Check(report.Suites[0].Cases[0].Stages, HasLen, 2)

}

func (s *CaseSuite) TestRunnerReportWithIncorrectRequest(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_incorrect_baidu_request")

	// then
	c.Check(err, IsNil)

	// when
	_, err = test_report.GetReport().Gen("./gen")

	// then
	file, _ := os.ReadFile("./gen/testrunner-report.json")
	report := test_report.Report{}
	_ = json.Unmarshal([]byte(file), &report)
	c.Check(report, NotNil)
	c.Check(report.Pass, Equals, false)
	c.Check(report.Suites, HasLen, 1)
	c.Check(report.Suites[0].Pass, Equals, false)
	c.Check(report.Suites[0].Cases, HasLen, 2)
	c.Check(report.Suites[0].Cases[0].RunTime, NotNil)
	c.Check(report.Suites[0].Cases[0].Pass, Equals, false)
	c.Check(report.Suites[0].Cases[0].Stages, HasLen, 2)

}

//todo: finish the unit-test case
//func (s *CaseSuite) TestRunnerReportWithNilAssertion(c *C) {
//	// given
//	testrunner := test_runner.Testrunner{}
//
//	// when
//	err := testrunner.Run("./data/testsuite_nil_assertion")
//
//	// then
//	c.Check(err, IsNil)
//
//	// when
//	_, err = test_report.GetReport().Gen("./gen")
//
//	// then
//	file, _ := os.ReadFile("./gen/testrunner-report.json")
//	report := test_report.Report{}
//	_ = json.Unmarshal([]byte(file), &report)
//	c.Check(report, NotNil)
//	c.Check(report.Pass, Equals, false)
//	c.Check(report.Suites, HasLen, 1)
//	c.Check(report.Suites[0].Pass, Equals, false)
//	c.Check(report.Suites[0].Cases, HasLen, 2)
//	c.Check(report.Suites[0].Cases[0].RunTime, NotNil)
//	c.Check(report.Suites[0].Cases[0].Pass, Equals, false)
//	c.Check(report.Suites[0].Cases[0].Stages, HasLen, 2)
//
//}
