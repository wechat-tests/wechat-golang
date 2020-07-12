package wxxx

import (
	"github.com/pkg/errors"
)

func AssertNotNil(o interface{}, msg string) {
	if o != nil {
		return
	}
	panic(errors.Errorf("assert not nil failed ! %s", msg))
}

func AssertTrue(cond bool, msg string) {
	if cond {
		return
	}
	panic(errors.Errorf("assert must be true failed ! %s", msg))
}

func AssertFalse(cond bool, msg string) {
	if !cond {
		return
	}
	panic(errors.Errorf("assert must be false failed ! %s", msg))
}

func AssertNotEmpty(str string, msg string) {
	if IsNotEmptyStr(str) {
		return
	}
	panic(errors.Errorf("assert must not empty failed ! %s", msg))
}

func AssertError(err error, msg string) {
	if err == nil {
		return
	}
	panic(errors.Wrapf(err, "%s :causedBy", msg))
}
