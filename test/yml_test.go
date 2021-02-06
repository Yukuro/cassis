package test

import (
	"cli/src/agent-docker-compose"
	"cli/src/commander"
	"cli/src/ledger-docker-compose"
	"fmt"
	"path/filepath"
	"testing"
)

func TestRenameNetworks(t *testing.T) {
	ledger_docker_compose.RenameNetworks("/home/tuple/GolandProjects/cli/test/", "shared-network01")
}

func TestGenerateDockerCompose(t *testing.T) {
	//agent.ConvertFromGraph("/home/tuple/GolandProjects/cli/test/test.dot")
}

//func TestSimpleRead(t *testing.T) {
//	agent.SimpleRead("/home/tuple/GolandProjects/cli/common/docker-compose.yml")
//}

//func TestConvert(t *testing.T) {
//	workdir := "/home/tuple/GolandProjects/cli/test/"
//	repoName := "aries-cloudagent-python"
//	if commander.IsExistDir(filepath.Join(workdir, repoName)) == false {
//		agent.ConvertFromGraph(
//			"/home/tuple/GolandProjects/cli/test/test.dot",
//			"/home/tuple/GolandProjects/cli/sandbox",
//			"shared-exp-net",
//			"192.168.3.16",
//		)
//	}
//}

func TestGetServiceName(t *testing.T) {
	workdir := "/home/tuple/GolandProjects/cli/sandbox/"
	_, err := agent_docker_compose.GetServiceNameFromWorkdir(workdir)
	if err != nil {
		panic(err)
	}
}

func TestIsExist(t *testing.T) {
	workdir := "/home/tuple/GolandProjects/cli/test/"
	repoAgent := "aries-cloudagent-python"
	repoVon := "von-network"
	fmt.Println(commander.IsExistDir(filepath.Join(workdir, repoAgent)))
	fmt.Println(commander.IsExistDir(filepath.Join(workdir, repoVon)))
}
