package container

import (
	"os/exec"
)

type Image struct {
	name string
}

func CreateImage(name, path string) (*Image, error) {
	cmd := exec.Command("docker", "build", "--no-cache", "--rm", "-t", name, path)
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return &Image{name: name}, nil
}

func (i *Image) NewContainer(name string) *Container {
	return New(i.name, name)
}

func (i *Image) Remove() error {
	cmd := exec.Command("docker", "rmi", i.name)
	return cmd.Run()
}
