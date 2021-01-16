package test

import (
	"cli/src/commander"
	"fmt"
	"testing"
)

//func TestGetSeed(t *testing.T){
//	agentNameList := []string{
//		"Issuer1",
//		"Issuer2",
//		"Holder1",
//		"Holder2",
//	}
//
//	fmt.Println(commander.RegisterDid(agentNameList))
//}

func TestComLedger(t *testing.T){
	alias := "Issuer123"
	seed := "mGbYfizQPdrrxiPcNaKkfHdDBKboPEhA"
	publicDid, err := commander.ComLedger(alias, seed)
	if err != nil{
		panic(err)
	}
	fmt.Println(publicDid)
}

func TestRegisterDid(t *testing.T){
	agentNameList := []string{
		"Issuer_fRGAu",
		"Issuer_wZCDm",
		"Holder_ifMXo",
		"Holder_sikDE",
		"Verifier_WkvFN",
	}

	agentAndDid, err := commander.RegisterDID(agentNameList)
	if err != nil{
		panic(err)
	}

	fmt.Println(agentAndDid)
}
