package main

import (
	"fmt"
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
	img, err := container.CreateImage("remexec/test", imgPath)
	if err != nil {
		fmt.Println("Could not build docker image.")
		return 1
	}
	defer img.Remove()

	cnt := img.NewContainer("remote")
	if err := cnt.StartAndPublish("12345:22"); err != nil {
		fmt.Println("Could not start docker container.")
		return 1
	}
	defer cnt.Terminate()

	return m.Run()
}
