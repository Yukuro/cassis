package commander

import (
	"errors"
	"github.com/creack/pty"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildVonNetwork(workdir string) error {
	workdir = filepath.Join(workdir, "von-network")
	os.Chdir(workdir)
	cmd := exec.Command("./manage", "build")
	fl, err := pty.Start(cmd)
	if err != nil {
		return errors.New("can't build von-network")
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func StartVonNetwork(workdir string) error {
	workdir = filepath.Join(workdir, "von-network")
	os.Chdir(workdir)
	cmd := exec.Command("./manage", "start")
	fl, err := pty.Start(cmd)
	if err != nil {
		return errors.New("can't start von-network")
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func GetMyIPAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			return string(ip), nil
		}
	}
	return "", nil
}

func IsExistDir(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// TODO ひとつにまとめる validateIsDotfile
func IsExistDotInDir(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, files := range files {
		if filepath.Ext(files.Name()) == ".dot" {
			return true, nil
		}
	}
	return false, nil
}

func RemoveAllFilesAtWorkDir(workdir string) error {
	parameter := []string{
		"-rf",
		"./*",
	}

	os.Chdir(workdir)
	cmd := exec.Command("rm", parameter...)
	fl, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func MapKeyToSlice(source map[string]string) []string {
	var result []string
	for key, _ := range source {
		result = append(result, key)
	}
	return result
}

// ngrok
func ExposeNgrok8001To8006(workdir string) error {
	parameter := []string{
		"ngrok",
		"start",
		"-config",
		"/home/tuple/GolandProjects/cli/config-ngrok.yml",
		"issuer1",
		"holder1",
		"verifier1",
	}

	os.Chdir(workdir)
	err := exec.Command("nohup", parameter...).Start()
	if err != nil {
		return err
	}
	return nil
}
