package unit_test

import (
	"encoding/json"
	"os"
	"testing"
	"testrunner/common"
	"testrunner/test_report"
	"testrunner/test_runner"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RunnerSuite struct{}

var _ = Suite(&RunnerSuite{})

func (s *RunnerSuite) SetUpTest(c *C) {
	_ = os.Setenv(common.PhaseEnv, string(common.Asserting))
}

//func (s *RunnerSuite) TestRunnerCorrectCheck(c *C) {
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
//func (s *RunnerSuite) TestRunnerErrorCheck(c *C) {
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

func (s *RunnerSuite) TestRunnerRun(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_correct_import")

	// then
	c.Check(err, IsNil)
}

func (s *RunnerSuite) TestRunnerExec(c *C) {
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

func (s *RunnerSuite) TestRunnerReportWithCorrectRequest(c *C) {
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

func (s *RunnerSuite) TestRunnerReportWithIncorrectRequest(c *C) {
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
//func (s *RunnerSuite) TestRunnerReportWithNilAssertion(c *C) {
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

func (s *RunnerSuite) TestRunnerWithMultiLevelVars(c *C) {
	// given
	testrunner := test_runner.Testrunner{}
	// when
	err := testrunner.Run("./data/testsuite_multi_level_vars/only_testcase")
	// then
	c.Check(err, IsNil)
	// when
	_, err = test_report.GetReport().Gen("./gen/multi_level_vars/only_testcase")
	// then
	file, _ := os.ReadFile("./gen/testrunner-report.json")
	report := make(map[string]any, 0)
	_ = json.Unmarshal(file, &report)
	c.Check(report, NotNil)

	// given
	err = testrunner.Run("./data/testsuite_multi_level_vars/only_description")
	// then
	c.Check(err, IsNil)
	// when
	_, err = test_report.GetReport().Gen("./gen/multi_level_vars/only_description")
	// then
	file, _ = os.ReadFile("./gen/testrunner-report.json")
	report = make(map[string]any, 0)
	_ = json.Unmarshal(file, &report)
	c.Check(report, NotNil)

	// given
	err = testrunner.Run("./data/testsuite_multi_level_vars/both_desc_testcase")
	// then
	c.Check(err, IsNil)
	// when
	_, err = test_report.GetReport().Gen("./gen/multi_level_vars/both_desc_testcase")
	// then
	file, _ = os.ReadFile("./gen/testrunner-report.json")
	report = make(map[string]any, 0)
	_ = json.Unmarshal(file, &report)
	c.Check(report, NotNil)
}

func (s *RunnerSuite) TestRunnerWithRecordMode(c *C) {
	// given
	_ = os.Setenv(common.PhaseEnv, string(common.Recording))
	testrunner := test_runner.Testrunner{}
	// when
	err := testrunner.Run("./data/testsuite_ready_for_record")
	// then
	c.Check(err, IsNil)
}
