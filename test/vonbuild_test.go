package test

import (
	"cli/src/commander"
	"testing"
)

func TestBuildVon(t *testing.T) {
	commander.BuildVonNetwork("/home/tuple/GolandProjects/cli/test/")
}

func TestStartVon(t *testing.T){
	commander.StartVonNetwork("/home/tuple/GolandProjects/cli/test/")
}
