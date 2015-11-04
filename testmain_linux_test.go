package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"text/template"

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

	return m.Run()
}

type ConfigParam struct {
	Host   string
	GoPath string
}

func prepareTestData(remoteHost string) error {
	dataDir := filepath.Join(os.Getenv("GOPATH"), "src",
		"github.com", "unirita", "remexec-test", "_testdata")
	testBase := filepath.Join(os.Getenv("GOPATH"), "bin")

	param := &ConfigParam{Host: remoteHost, GoPath: os.Getenv("GOPATH")}
	if err := createConf("command.ini", dataDir, testBase, param); err != nil {
		return err
	}
	if err := createConf("script.ini", dataDir, testBase, param); err != nil {
		return err
	}
	if err := copyFile("remote.pem", dataDir, testBase); err != nil {
		return err
	}

	return nil
}

func createConf(name, fromDir, toDir string, param *ConfigParam) error {
	tpl, err := template.ParseFiles(filepath.Join(fromDir, name))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(toDir, name))
	if err != nil {
		return err
	}
	defer file.Close()

	return tpl.Execute(file, param)
}

func copyFile(name, from, to string) error {
	srcPath := filepath.Join(from, name)
	targetPath := filepath.Join(to, name)

	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	r := bufio.NewReader(src)
	w := bufio.NewWriter(target)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
	return w.Flush()
}
