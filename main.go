package main

import (
	"fmt"

	"github.com/ory-am/fosite"
)

//go:generate counterfeiter . AccessRequester
type AccessRequester interface {
	fosite.AccessRequester
}

func main() {
	fmt.Println("vim-go")
}
