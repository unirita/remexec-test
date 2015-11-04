package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/unirita/remexec-test/container"
)

func TestMain(m *testing.M) {
	os.Exit(realTestMain(m))
}

func realTestMain(m *testing.M) int {
	imgPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita",
		"remexec-test", "_docker", "remote")
	img := container.NewImage("remexec/test", imgPath)
	defer img.Remove()
}
