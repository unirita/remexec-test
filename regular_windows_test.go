package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unirita/remexec-test/capturer"
)

func TestRemote_Command_NoParameter(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand("pwd")
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 0 {
		t.Errorf("RC => %d, wants %d", rc, 0)
	}
	if !strings.Contains(output, "Path") {
		t.Errorf("Output is not expected format.")
		t.Errorf("Output: %s", output)
	}
}

func TestRemote_Command_WithParameter(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand(`echo testmessage`)
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 0 {
		t.Errorf("RC => %d, wants %d", rc, 0)
	}
	output = strings.Trim(output, "\r\n")
	if output != "testmessage" {
		t.Errorf("Output => %s, wants %s.", output, "testmessage")
	}
}

func TestRemote_Script(t *testing.T) {
	scriptPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com",
		"unirita", "remexec-test", "_testdata", "windows", "test.ps1")

	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand(scriptPath + ` test1 test2`)
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 12 {
		t.Errorf("RC => %d, wants %d", rc, 12)
	}
	if !strings.Contains(output, "test1") {
		t.Errorf("Output does not contains correct first parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
	if !strings.Contains(output, "test2") {
		t.Errorf("Output does not contains correct second parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
}

func TestLocalScript(t *testing.T) {
	scriptPath := filepath.Join(baseDir, "localtest.ps1")

	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeLocalScript(scriptPath + ` test1 test2`)
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}

	// NOTICE:
	// remexec can not get RC of local script.
	if rc != 0 {
		t.Errorf("RC => %d, wants %d", rc, 0)
	}

	if !strings.Contains(output, "test1") {
		t.Errorf("Output does not contains correct first parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
	if !strings.Contains(output, "test2") {
		t.Errorf("Output does not contains correct second parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
}
