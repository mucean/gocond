package gocond

import (
	"math/rand"
	"sync"
)

type Checker func(*Context) bool

func EmptyCtxChecker(chk func() bool) Checker {
	return func(c *Context) bool {
		return chk()
	}
}

type Conditional interface {
	Check(*Context) bool
}

type Needer interface {
	Need() bool
	Default() bool
}

type Condition struct {
	chk Checker
}

func (c *Condition) Check(ctx *Context) bool {
	return c.chk(ctx)
}

type NeedCond struct {
	Chk     Checker
	DftNeed bool
}

func (nc *NeedCond) DoCheck(c Needer, ctx *Context) bool {
	if c.Need() {
		return nc.Chk(ctx)
	}
	return c.Default()
}

func (nc *NeedCond) Default() bool {
	return nc.DftNeed
}

type RandCond struct {
	NeedCond
	MaxInt int
}

func (rc *RandCond) Need() bool {
	return rand.Intn(rc.MaxInt) == 0
}

func (rc *RandCond) Check(ctx *Context) bool {
	return rc.DoCheck(rc, ctx)
}

func NewNextRandCond(chk Checker, maxInt int, def bool) *NextRandCond {
	return &NextRandCond{
		RandCond: RandCond{
			NeedCond: NeedCond{
				Chk:     chk,
				DftNeed: def,
			},
			MaxInt: maxInt,
		},
	}
}

type NextRandCond struct {
	locker sync.RWMutex
	RandCond
	next bool
}

func (nc *NextRandCond) Check(ctx *Context) bool {
	res := nc.DoCheck(nc, ctx)
	nc.locker.Lock()
	if !nc.next && res != nc.Default() {
		nc.next = true
	}
	nc.locker.Unlock()
	return res
}

func (nc *NextRandCond) Need() bool {
	nc.locker.RLock()
	next := nc.next
	nc.locker.RUnlock()
	res := next
	if !next {
		res = nc.RandCond.Need()
	}
	return res
}
