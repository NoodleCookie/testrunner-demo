package unit_test

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"testrunner/test_assertion"
)

func Test(t *testing.T) { TestingT(t) }

type CaseSuite struct{}

var _ = Suite(&CaseSuite{})

func (s *CaseSuite) TestGenericAssertorEqual(c *C) {
	// given
	assertor := test_assertion.GenericAssertor{}

	// when
	assert, err := assertor.Assert(1, test_assertion.Equals, 1)

	// then
	c.Check(err, IsNil)
	c.Check(assert, Equals, true)

	// when
	assert, err = assertor.Assert("1", test_assertion.Equals, 1)

	// then
	c.Check(err, NotNil)
	c.Check(assert, Equals, false)

	// when
	assert, err = assertor.Assert("yes", test_assertion.Equals, "no")

	// then
	c.Check(err, NotNil)
	c.Check(assert, Equals, false)

	// when
	assert, err = assertor.Assert("yes", test_assertion.Equals, "yes")

	// then
	c.Check(err, IsNil)
	c.Check(assert, Equals, true)
}

func (s *CaseSuite) TestGenericAssertorEqualWithBeforeAfterMethod(c *C) {
	// given
	assertor := test_assertion.GenericAssertor{}

	// when
	assertor.Before(func() {
		fmt.Println("assert before")
	})
	assertor.After(func(result bool, err error) {
		fmt.Printf("assert after %v\n", result)
	})
	assert, err := assertor.Assert(1, test_assertion.Equals, 1)

	// then
	c.Check(err, IsNil)
	c.Check(assert, Equals, true)
}
