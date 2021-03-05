package test

import (
	"cli/src/commander"
	ngrok_analyze "cli/src/ngrok-analyze"
	"fmt"
	"testing"
)

func TestNgrok8001To8006(t *testing.T) {
	workdir := "/home/tuple/GolandProjects/cli"

	err := commander.ExposeNgrok8001To8006(workdir)
	if err != nil {
		panic(err)
	}
}

func TestGetNgrokUrl(t *testing.T) {
	urlList, err := ngrok_analyze.GetNgrokUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(urlList)
}
