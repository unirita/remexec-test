package container

import (
	"fmt"
	"os/exec"
	"strings"
)

type Container struct {
	image     string
	name      string
	isRunning bool
}

func New(image string, name string) *Container {
	c := new(Container)
	c.image = image
	c.name = name
	c.isRunning = false
	return c
}

func (c *Container) Start() error {
	if c.isRunning {
		return fmt.Errorf("Container [%s] already running.", c.name)
	}
	cmd := exec.Command("docker", "run", "-itd", "--name", c.name, c.image)
	err := cmd.Run()
	if err != nil {
		return err
	}

	c.isRunning = true
	return nil
}

func (c *Container) StartAndPublish(publish string) error {
	if c.isRunning {
		return fmt.Errorf("Container [%s] already running.", c.name)
	}
	cmd := exec.Command("docker", "run", "-itd", "-p", publish, "--name", c.name, c.image)
	err := cmd.Run()
	if err != nil {
		return err
	}

	c.isRunning = true
	return nil
}

func (c *Container) IPAddress() (string, error) {
	out, err := exec.Command("docker", "inspect",
		"--format='{{.NetworkSettings.IPAddress}}'", c.name).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func (c *Container) Terminate() {
	if !c.isRunning {
		return
	}
	exec.Command("docker", "stop", c.name).Run()
	exec.Command("docker", "rm", "-f", c.name).Run()
	c.isRunning = false
}
