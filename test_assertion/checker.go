package test_assertion

import (
	"fmt"
	"github.com/pkg/errors"
)

var Equals isEqual

type isEqual struct {
}

func (i isEqual) Check(actual, expect interface{}) (result bool, error error) {
	if actual == expect {
		return true, nil
	}
	return false, errors.New(fmt.Sprintf("%T(%v) is not equals %T(%v)", actual, actual, expect, expect))
}
