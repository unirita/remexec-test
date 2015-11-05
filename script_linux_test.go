package main

import (
	"path/filepath"
	"regexp"
	"testing"

	"github.com/unirita/remexec-test/capturer"
)

func TestLocalScript(t *testing.T) {
	scriptPath := filepath.Join(baseDir, "localtest.sh")

	c := capturer.NewStdoutCapturer()
	c.Start()
	rc, err := executeLocalScript(scriptPath + ` "test1" "test2"`)
	output := c.Stop()

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 23 {
		t.Errorf("RC => %d, wants %d", rc, 23)
	}

	matcher := regexp.MustCompile(`script=/home/keyuser/tmp/\d{14}\.\d{3}\.localtest\.sh`)
	if !matcher.MatchString(output) {
		t.Errorf("Output does not contains correct script value.")
		t.Log("Output:")
		t.Log(output)
	}
	if !strings.Contains(output, "param1=test1") {
		t.Errorf("Output does not contains correct first parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
	if !strings.Contains(output, "param2=test2") {
		t.Errorf("Output does not contains correct second parameter value.")
		t.Log("Output:")
		t.Log(output)
	}
}
