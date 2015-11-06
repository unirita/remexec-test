package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type ConfigParam struct {
	Host   string
	GoPath string
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
