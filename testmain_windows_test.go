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

	prepareTestData()

	return m.Run()
}

func prepareTestData() error {
	dataDir := filepath.Join(os.Getenv("GOPATH"), "src",
		"github.com", "unirita", "remexec-test", "_testdata", "windows")

	param := new(ConfigParam)
	param.User = user
	param.Pass = pass

	if err := createConf("command.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("script.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("noexisthost.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := createConf("wrongpass.ini", dataDir, baseDir, param); err != nil {
		return err
	}
	if err := copyFile("localtest.sh", dataDir, baseDir); err != nil {
		return err
	}

	return nil
}
