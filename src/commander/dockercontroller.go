package commander

import (
	"errors"
	"os/exec"
)

func BuildDockerNetwork(networkName string)(string, error) {
	parameter := []string{
		"network",
		"create",
		networkName,
	}

	out, err := exec.Command("docker", parameter...).Output()
	if err != nil{
		return "", errors.New("can't execute docker network create")
	}
	return string(out), nil
}