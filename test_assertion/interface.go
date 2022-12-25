package test_assertion

type (
	Assertor interface {
		Assert(actual interface{}, checker Checker, expect interface{}) (bool, error)
	}
	Checker interface {
		Check(actual, expect interface{}) (result bool, error error)
	}
)

var (
	_ Checker  = (*isEqual)(nil)
	_ Checker  = (*isMatchRegex)(nil)
	_ Assertor = (*GenericAssertor)(nil)
	_ Assertor = (*JsonAssertor)(nil)
)
