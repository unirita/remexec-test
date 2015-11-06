package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unirita/remexec-test/capturer"
)

func TestIrregular_ConfigFileNotExists(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeWithConfig(filepath.Join(baseDir, "noexistconfig.ini"))
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}
	if !strings.Contains(output, "REX002E") {
		t.Errorf("Output does not contains expected error message.")
		t.Log("Output:")
		t.Log(output)
	}
}

func TestIrregular_HostNotExists(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeWithConfig(filepath.Join(baseDir, "noexisthost.ini"))
	c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}
}

func TestIrregular_WrongPassword(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeWithConfig(filepath.Join(baseDir, "wrongpass.ini"))
	c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}
}

func TestIrregular_RemoteCommandNotExists(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand("noexistcommand")
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}
	if strings.Contains(output, "REX003E") {
		t.Errorf("Output contains unexpected error message.")
		t.Log("Output:")
		t.Log(output)
	}
}

func TestIrregular_LocalScriptNotExists(t *testing.T) {
	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeLocalScript("noexistscript")
	c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 0 {
		t.Errorf("RC => %d, wants %d", rc, 0)
	}
}

func TestIrregular_RemexecPs1NotExists(t *testing.T) {
	scriptName := filepath.Join(baseDir, "remexec.ps1")
	tmpScriptName := filepath.Join(baseDir, "remexec.ps1.tmp")

	os.Rename(scriptName, tmpScriptName)
	defer os.Rename(tmpScriptName, scriptName)

	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeRemoteCommand("pwd")
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}

	if rc != 255 {
		t.Errorf("RC => %d, wants %d", rc, 255)
	}

	if strings.Contains(output, "REX003E") {
		t.Errorf("Output contains unexpected error message.")
		t.Log("Output:")
		t.Log(output)
	}
}
