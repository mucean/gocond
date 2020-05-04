package gocond

import (
	"fmt"
	"testing"
)

func Test_ABC(t *testing.T) {
	c := NewNextRandCond(func(c *Context) bool {
		println("aa")
		return false
	}, 2, false)
	for i := 0; i < 20; i++ {
		fmt.Println(c.Check(nil))
	}
}
