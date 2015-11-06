package main

import (
	"os"
	"path/filepath"
)

var baseDir = filepath.Join(os.Getenv("GOPATH"), "bin")

func TestMain(m *testing.M) {
	os.Exit(realTestMain(m))
}
