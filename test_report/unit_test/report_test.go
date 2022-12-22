package unit_test

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"testrunner/common"
	"testrunner/test_report"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

func (s *CaseSuite) TestReportGen(c *C) {
	// given
	report := test_report.Report{}

	// when
	gen, err := report.Gen("./gen")

	// then
	c.Check(err, IsNil)
	c.Check(gen, Equals, fmt.Sprintf("gen/%s.%s", common.ReportFile, common.ReportFileExt))
}

func (s *CaseSuite) TestReportAppendStageAndGen(c *C) {
	// given
	report := test_report.Report{}

	// when
	report.AppendSuite("first_suite")
	report.AppendCase("first_case")
	report.AppendStage("first_stage", true, "I'm pass")
	gen, err := report.Gen("./gen/success")

	// then
	c.Check(err, IsNil)
	c.Check(gen, Equals, fmt.Sprintf("gen/success/%s.%s", common.ReportFile, common.ReportFileExt))
	c.Check(report.Pass, Equals, true)

	// when
	report.AppendStage("second_stage", false, "I was failed")
	c.Check(report.Pass, Equals, true)
	gen, err = report.Gen("./gen/failed")

	// then
	c.Check(err, IsNil)
	c.Check(gen, Equals, fmt.Sprintf("gen/failed/%s.%s", common.ReportFile, common.ReportFileExt))
	c.Check(report.Pass, Equals, false)
}
