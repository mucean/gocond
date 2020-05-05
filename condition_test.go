package gocond

import (
	"testing"
)

var nilChecker Checker = nil
var trueChecker Checker = func(ctx *Context) bool {
	return true
}
var falseChecker Checker = func(ctx *Context) bool {
	return false
}
var maxInt = 5
var mustRunMaxInt = 1
var dftRes = true

func Test_NewNeedCond_MustCheck(t *testing.T) {
	checked := false
	c := NewRandCond(func(ctx *Context) bool {
		checked = true
		return true
	}, mustRunMaxInt, dftRes)
	c.Check(nil)
	if !checked {
		t.Errorf("Check must run")
		t.FailNow()
	}
}

func Test_NewNeedCond_NotCheck(t *testing.T) {
	checked := false
	c := NewNeedCond(func(ctx *Context) bool {
		checked = true
		return true
	}, NoNeed{dft: false})
	c.Check(nil)
	if checked {
		t.Errorf("Check must not run")
		t.FailNow()
	}
}

func Test_NewNextNeedCond_NoNeed_DFT_False(t *testing.T) {
	checked := false
	checkRes := true
	noNeed := NoNeed{dft: checkRes}
	c := NewNextNeedCond(func(ctx *Context) bool {
		checked = true
		return checkRes
	}, noNeed)

	if res := c.Check(nil); res != checkRes {
		t.Errorf("left is: %v, right is: %v", res, checkRes)
		t.FailNow()
	}
	if !checked {
		t.Errorf("Checker must run")
		t.FailNow()
	}

	if c.need.Default() == checkRes {
		if c.next {
			t.Errorf("next property must be false")
			t.FailNow()
		}
	}

	res := c.Check(nil)
	if res != c.need.Default() {
		t.Errorf("res is: %v, default is: %v", res, c.need.Default())
		t.FailNow()
	} else {
		if c.next {
			t.Errorf("next property must be false")
			t.FailNow()
		}
	}

	checked = false
}

func Test_NewNextNeedCond_NoNeed_DFT_True(t *testing.T) {
	checked := false
	checkRes := true
	noNeed := NoNeed{dft: false}

	c := NewNextNeedCond(func(ctx *Context) bool {
		checked = true
		return checkRes
	}, noNeed)

	if res := c.Check(nil); res != checkRes {
		t.Errorf("left is: %v, right is: %v", res, checkRes)
		t.FailNow()
	}

	if c.need.Default() != checkRes {
		if !c.next {
			t.Errorf("next property must be true")
			t.FailNow()
		}
	}

	if !checked {
		t.Errorf("Checker must run")
		t.FailNow()
	}
}

func Test_NewNextNeedCond_AlwayNeed_DFT_True(t *testing.T) {
	checked := false
	checkRes := true
	alwayNeed := NewRandNeed(1, true)

	c := NewNextNeedCond(func(ctx *Context) bool {
		checked = true
		return checkRes
	}, alwayNeed)

	if res := c.Check(nil); res != checkRes {
		t.Errorf("left is: %v, right is: %v", res, checkRes)
		t.FailNow()
	}

	if c.need.Default() == checkRes {
		if c.next {
			t.Errorf("next property must be false")
			t.FailNow()
		}
	}

	if !checked {
		t.Errorf("Checker must run")
		t.FailNow()
	}

	checkRes = false
	if res := c.Check(nil); res != checkRes {
		t.Errorf("left is: %v, right is: %v", res, checkRes)
		t.FailNow()
	}
	if c.need.Default() != checkRes {
		if !c.next {
			t.Errorf("next property must be true")
			t.FailNow()
		}
	}
}
