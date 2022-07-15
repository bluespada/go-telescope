package main

import (
	"github.com/bluespada/go-telescope/server"
)

func main() {
	telescope := new(server.TelescopeServer)
	telescope.RunEngine()
}
