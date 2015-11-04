package main

import (
	"strings"
	"testing"

	"github.com/unirita/remexec-test/capturer"
)

func TestCommand_NoParameter(t *testing.T) {
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
	if !strings.HasPrefix(output, "/") {
		t.Errorf("Output is not expected format.")
		t.Errorf("Output: %s", output)
	}
}

func TestCommand_WithParameter(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand(`echo "testmessage"`)
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 0 {
		t.Errorf("RC => %d, wants %d", rc, 0)
	}
	if !output != "testmessage" {
		t.Errorf("Output => %s, wants %s.", output, "testmessage")
	}
}
