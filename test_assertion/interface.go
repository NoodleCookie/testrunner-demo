package test_assertion

type (
	Assertor interface {
		Assert(actual interface{}, checker Checker, expect interface{}) (bool, error)
	}

	AbstractAssertor struct {
		before func()
		after  func(result bool, err error)
	}

	Checker interface {
		Check(actual, expect interface{}) (result bool, error error)
	}
)

var (
	_ Assertor = (*AbstractAssertor)(nil)
	_ Checker  = (*isEqual)(nil)
)

func (a *AbstractAssertor) Assert(actual interface{}, checker Checker, expect interface{}) (bool, error) {
	if a.before != nil {
		a.before()
	}
	
	result, err := checker.Check(actual, expect)

	if a.after != nil {
		a.after(result, err)
	}

	return result, err
}

func (a *AbstractAssertor) Before(before func()) {
	a.before = before
}

func (a *AbstractAssertor) After(after func(result bool, err error)) {
	a.after = after
}
