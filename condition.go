package gocond

import (
	"sync"
)

// Checker is a type name that wrap a function take a Context pointer
// and return a boolean
type Checker func(*Context) bool

// Conditional is a interface representing something is true
type Conditional interface {
	// Check function will take a Context pointer for parameters and
	// return a boolean to caller
	Check(*Context) bool
}

// A NeedCond struct is connector that the checker can be call by the result of the
// Need method
type NeedCond struct {
	need Needer
	// the checker function
	chk Checker
}

func (nc *NeedCond) Check(ctx *Context) bool {
	if nc.need.Need(ctx) {
		return nc.chk(ctx)
	}
	return nc.need.Default()
}

// NewNeedCond function initializes NeedCond
func NewNeedCond(chk Checker, need Needer) *NeedCond {
	return &NeedCond{
		chk:  chk,
		need: need,
	}
}

// NewRandCond function initialize NeedCond with a RandNeed Needer
func NewRandCond(chk Checker, maxInt int, dft bool) *NeedCond {
	return NewNeedCond(chk, NewRandNeed(maxInt, dft))
}

// NewNextRandCond initializes the NextNeedCond struct with the checker function
// and a Needer
func NewNextNeedCond(chk Checker, need Needer) *NextNeedCond {
	return &NextNeedCond{
		chk:  chk,
		need: need,
		next: true,
	}
}

// The difference between NextNeedCond and NeedCond is when the Checker returns
// not equal de Needer's default value, Check method must call in the next called
type NextNeedCond struct {
	locker sync.RWMutex
	chk    Checker
	need   Needer
	next   bool
}

func (nc *NextNeedCond) Check(ctx *Context) bool {
	nc.locker.RLock()
	next := nc.next
	nc.locker.RUnlock()
	run := next
	if !next {
		run = nc.need.Need(ctx)
	}
	dft := nc.need.Default()
	res := dft
	if run {
		res = nc.chk(ctx)
		nc.locker.Lock()
		if nc.next {
			if res == dft {
				nc.next = false
			}
		} else {
			if res != dft {
				nc.next = true
			}
		}
		nc.locker.Unlock()
	} else {
		nc.locker.Lock()
		if nc.next {
			nc.next = false
		}
		nc.locker.Unlock()
	}
	return res
}
