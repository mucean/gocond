// Package gocond provides ...
package gocond

import "fmt"

type Runner struct {
}

func (r *Runner) Run(f func() error) {
	fmt.Println("hello")
}

// TODO
