package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/unirita/remexec-test/capturer"
)

func TestIrregular_HostNotExists(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeWithConfig(filepath.Join(baseDir, "noexisthost.ini"))
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}
	if !strings.Contains(output, "REX003E") {
		t.Errorf("Output does not contains expected error message.")
		t.Log("Output:")
		t.Log(output)
	}
}
