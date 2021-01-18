package commander

import (
	"errors"
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
)

func BuildDockerNetwork(networkName string) (string, error) {
	parameter := []string{
		"network",
		"create",
		networkName,
	}

	out, err := exec.Command("docker", parameter...).Output()
	if err != nil {
		return "", errors.New("can't execute docker network create")
	}
	return string(out), nil
}

func DockerComposeUpAtWorkdir(workdir string) error {
	parameter := []string{
		"up",
		"-d",
	}

	os.Chdir(workdir)
	cmd := exec.Command("docker-compose", parameter...)
	fl, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, fl)
	return nil
}
