package models

import "fmt"

var DEBUG bool

func init() {
	DEBUG = false
}

func debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format+"\n", args...)
	}
}
