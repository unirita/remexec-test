package main

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

var (
	remexec    = filepath.Join(baseDir, "remexec")
	commandIni = filepath.Join(baseDir, "command.ini")
	scriptIni  = filepath.Join(baseDir, "script.ini")
)

func executeRemoteCommand(command string) (int, error) {
	cmd := exec.Command(remexec, "-c", commandIni, "-e", command)
	return getRC(cmd.Run())
}

func getRC(err error) (int, error) {
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				return s.ExitStatus(), nil
			}
		}
		return -1, err
	}
	return 0, nil
}
