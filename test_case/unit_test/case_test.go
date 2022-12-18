package unit_test

import (
	"gopkg.in/yaml.v2"
	"os"
	"testing"
	"testrunner/test_case"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

func (s *CaseSuite) TestUnmarshalCaseFile(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_case.Case{}

	// when
	err := yaml.Unmarshal(file, tc)

	// then
	c.Check(err, IsNil)
	c.Check(tc.Stages, HasLen, 2)
	c.Check(tc.Stages[0].Request.Body, Equals, "{\"sequence\": \"hello world\"}")
}

func (s *CaseSuite) TestCaseExecute(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_case.Case{}

	// when
	_ = yaml.Unmarshal(file, tc)
	err := tc.Execute()

	// then
	c.Check(err, IsNil)
}

func (s *CaseSuite) TestCaseExecuteWithFilter(c *C) {
	// given
	file, _ := os.ReadFile("./data/hello-world-test.yaml")
	tc := &test_case.Case{}

	// when
	_ = yaml.Unmarshal(file, tc)
	tc.AddFilter(func(stage test_case.Stage) bool {
		return stage.Mold == test_case.API
	})
	err := tc.Execute()

	// then
	c.Check(err, IsNil)
}
