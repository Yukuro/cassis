package commander

import (
	"errors"
	"fmt"
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func RemoveAllDockerImages() error {
	parameter := []string{
		"rmi",
		"$(docker",
		"images",
		"-q)",
	}

	cmd := exec.Command("docker", parameter...)
	fl, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func GetAdminPortFromWorkdir(workdir string) (map[string]string, error) {
	//d := filepath.Dir(workdir)
	dirName := filepath.Base(workdir)

	nodeName := []string{
		"issuer1",
		"issuer2",
		"issuer3",
		"issuer4",
		"issuer5",
		"holder1",
		"holder2",
		"holder3",
		"holder4",
		"holder5",
		"verifier1",
		"verifier2",
		"verifier3",
		"verifier4",
		"verifier5",
		"custom1",
		"custom2",
		"custom3",
		"custom4",
		"custom5",
	}

	nodeNameAndPort := map[string]string{}

	for _, node := range nodeName {
		containerName := fmt.Sprintf("%v_%v_1", dirName, node)

		parameter := []string{
			"port",
			containerName,
			"11000/tcp",
		}

		out, err := exec.Command("docker", parameter...).Output()

		if err != nil {
			continue
		}

		//permit := fmt.Sprintf("Error: No such container: %v\n", containerName)
		//if err != nil{
		//	//fmt.Println(err)
		//
		//	//TODO No such containerをキャッチする
		//	//if permit == string(os.Stderr)
		//}
		//fmt.Println(out)

		p := string(out)
		pe := strings.Replace(p, "0.0.0.0:", "", -1)
		pee := strings.Replace(pe, "\n", "", -1)
		k := strings.ToUpper(node[:1]) + node[1:]
		nodeNameAndPort[k] = pee
	}
	return nodeNameAndPort, nil
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

func DockerComposeDownAtWorkdir(workdir string) error {
	os.Chdir(workdir)
	cmd := exec.Command("docker-compose", "down")
	fl, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func DockerComposeDownAtVonNw(workdir string) error {
	os.Chdir(filepath.Join(workdir, "von-network"))
	cmd := exec.Command("./manage", "down")
	fl, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, fl)
	return nil
}
