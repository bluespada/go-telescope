package server

import (
	"fmt"
)

var (
	major   = 1
	minor   = 0
	patch   = 0
	build   = "dev"
	Version = fmt.Sprint(major, ".", minor, ".", patch, "-", build)
)
