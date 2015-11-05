package main

import (
	"os"
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return getRC(cmd.Run())
}

func executeLocalScript(script string) (int, error) {
	cmd := exec.Command(remexec, "-c", scriptIni, "-f", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return getRC(cmd.Run())
}

func executeWithConfig(config string) (int, error) {
	cmd := exec.Command(remexec, "-c", config, "-e", "pwd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
