package docker

import (
	"cli/src/commander"
	"fmt"
	"testing"
)

func TestBuildDockerNetwork(t *testing.T) {
	networkName := "shared-nw003"
	out, err := commander.BuildDockerNetwork(networkName)
	if err != nil{
		panic(err)
	}
	fmt.Printf("created %v : %v", networkName, out)
}
