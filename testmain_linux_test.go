package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/unirita/remexec-test/container"
)

func realTestMain(m *testing.M) int {
	imgPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "unirita",
		"remexec-test", "_docker", "remote")
	img, err := container.CreateImage("remexec/test", imgPath)
	if err != nil {
		fmt.Println("Error: Could not build docker image.")
		fmt.Println(err)
		return 1
	}
	defer img.Remove()

	cnt := img.NewContainer("remote")
	if err := cnt.StartAndPublish("12345:22"); err != nil {
		fmt.Println("Error: Could not start docker container.")
		fmt.Println(err)
		return 1
	}
	defer cnt.Terminate()

	remoteHost, err := cnt.IPAddress()
	if err != nil {
		fmt.Println("Error: Could not get container IP.")
		fmt.Println(err)
		return 1
	}

	if err := prepareTestData(remoteHost); err != nil {
		fmt.Println("Error: Could not prepare test data.")
		fmt.Println(err)
		return 1
	}

	time.Sleep(5 * time.Second)

	return m.Run()
}

func prepareTestData(remoteHost string) error {
	dataDir := filepath.Join(os.Getenv("GOPATH"), "src",
		"github.com", "unirita", "remexec-test", "_testdata", "linux")

	param := &ConfigParam{Host: remoteHost, GoPath: os.Getenv("GOPATH")}
	if err := createConf("command.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("script.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("noexisthost.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("noexistkey.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("wrongkey.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("wrongpass.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := copyFile("localtest.sh", dataDir, baseDir); err != nil {
		return err
	}
	if err := copyFile("remote.pem", dataDir, baseDir); err != nil {
		return err
	}
	if err := copyFile("wrong.pem", dataDir, baseDir); err != nil {
		return err
	}

	os.Chmod(filepath.Join(baseDir, "remote.pem"), 0400)

	return nil
}
