package main

import (
	"os"
	"path/filepath"
	"testing"
)

func realTestMain(m *testing.M) int {
	ps1Dir := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita",
		"remexec", "script")
	copyFile("remexec.ps1", ps1Dir, baseDir)

	return m.Run()
}
