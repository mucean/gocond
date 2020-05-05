package gocond

import "math/rand"

// Needer is the interface that wraps Need and Default method
type Needer interface {
	// Need returns a boolean to caller for if the Conditional Check method
	// need to call
	Need(*Context) bool
	// Default is a method that when the Need method return false, the Check method
	// will not execute and the default result will be
	Default() bool
}

func NewRandNeed(maxInt int, dft bool) *RandNeed {
	return &RandNeed{
		maxInt: maxInt,
		dft:    dft,
	}
}

type RandNeed struct {
	maxInt int
	dft    bool
}

func (r *RandNeed) Need(ctx *Context) bool {
	return rand.Intn(r.maxInt) == 0
}

func (r *RandNeed) Default() bool {
	return r.dft
}

func NewNoNeed(dft bool) NoNeed {
	return NoNeed{dft: dft}
}

type NoNeed struct {
	dft bool
}

func (n NoNeed) Need(ctx *Context) bool {
	return false
}

func (n NoNeed) Default() bool {
	return n.dft
}
