package unit_test

import (
	"testing"
	"testrunner/test_runner"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

func (s *CaseSuite) TestRunnerCorrectCheck(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	description, err := testrunner.CheckDescription("./data/testsuite_correct_import")

	// then
	c.Check(err, IsNil)
	c.Check(description.Import, HasLen, 1)
	c.Check(description.Import[0], Equals, "hello-test")
}

func (s *CaseSuite) TestRunnerErrorCheck(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	_, err := testrunner.CheckDescription("./data/testsuite_nil_import")

	// then
	c.Check(err, NotNil)
	c.Assert(err.Error(), Equals, "you must import your testcase into description")
}

func (s *CaseSuite) TestRunnerRun(c *C) {
	// given
	testrunner := test_runner.Testrunner{}

	// when
	err := testrunner.Run("./data/testsuite_correct_import")

	// then
	c.Check(err, IsNil)
}
