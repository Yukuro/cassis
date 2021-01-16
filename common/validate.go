package common

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func validateIsDotfile(input string) error{
	_, err := os.Stat(input)
	if os.IsNotExist(err){
		return errors.New("the file is not found")
	}

	if filepath.Ext(input) != ".dot" {
		return errors.New("the file is not dot file")
	}
	return nil
}

func validateIsCsvfile(input string) error{
	_, err := os.Stat(input)
	if os.IsNotExist(err){
		return errors.New("the file is not found")
	}

	if filepath.Ext(input) != ".csv" {
		return errors.New("the file is not csv file")
	}
	return nil
}

func validateExistFileAndDirectory(input string) error {
	_, err := os.Stat(input)
	if os.IsNotExist(err){
		return errors.New("the file is not found")
	}
	return nil
}

//TODO
func validateExistDockerNetwork(input string) error {
	parameter := []string{
		"network",
		"ls",
	}

	out, err := exec.Command("docker", parameter...).Output()
	if err != nil{
		return errors.New("network name is duplicated")
	}
	fmt.Println(out)
	return nil
}
