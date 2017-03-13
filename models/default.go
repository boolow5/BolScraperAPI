package models

import "fmt"

var DEBUG bool

func init() {
  fmt.Println("Initializing models")
  DEBUG = true
}

func debug(format string, args ...interface{}) {
  fmt.Printf(format + "\n", args...)
}
