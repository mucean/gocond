package gocond

import (
	"github.com/mucean/gotools/store"
)

// NewContext return a Context struct pointer, and initialized Args and Ret attributes
func NewContext() *Context {
	return &Context{Args: store.New(), Ret: store.New()}
}

// A Context is a container that contain arguments passed by user, and response by function, also
// a error when error occurred
type Context struct {
	Args store.Store
	Ret  store.Store
}
