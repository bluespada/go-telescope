package server_test

import (
	"testing"

	"github.com/bluespada/go-telescope/server"
)

func TestServerVersion(t *testing.T) {
	// Check server version is same or nah
	if server.Version == "1.0.0-dev" {
		println(server.Version)
	} else {
		t.FailNow()
	}
}

func TestServerS(t *testing.T) {
	// check some function here
	testServer := new(server.TelescopeServer)
	testServer.RunEngine()
}
