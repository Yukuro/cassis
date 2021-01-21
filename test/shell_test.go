package test

import (
	"cli/src/commander"
	"fmt"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	workdir := "/home/tuple/GolandProjects/cli/sandbox"
	err := commander.RemoveAllFilesAtWorkDir(workdir)
	if err != nil {
		panic(err)
	}
}

func TestGetAdminPort(t *testing.T) {
	workdir := "/home/tuple/GolandProjects/cli/sandbox"
	res, err := commander.GetAdminPortFromWorkdir(workdir)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
