package models

import (
	"fmt"
	"os"
	"strconv"
)

var DEBUG bool

func init() {
	DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
}

func debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format+"\n", args...)
	}
}
