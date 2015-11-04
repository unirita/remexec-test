package main

import (
	"testing"
)

func TestCommand_NoOption(t *testing.T) {
	rc, err := executeRemoteCommand("hostname")
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if rc != 0 {
		t.Errorf("RC => %s, wants %s", rc)
	}
}
