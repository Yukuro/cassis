/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SYSTEM REQUIREMENTS
install docker and add ENV
*/
package cmd

import (
	"cli/common"
	"cli/src/analyse"
	"cli/src/analyse/agent"
	"cli/src/analyse/ledger"
	"cli/src/commander"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildAgent()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func buildAgent() error {
	//get .dot file
	fmt.Println("Enter your dot file...")
	dotFilepath, err := common.PromptForDot("Dot")
	if err != nil {
		return err
	}

	analyse.Browse(dotFilepath)

	//Does Issuer - Holder =~ VC exist ?
	if analyse.RequireVC(dotFilepath) {
		fmt.Println("detected Issuer - Holder relation...")

		csvFilepath, err := common.PromptForCsv("Csv")
		if err != nil {
			return err
		}
		fmt.Printf("Your csv file is %v\n", csvFilepath)
	}

	fmt.Println("Start building Network and Agent...")

	workdir, err := common.PromptForFileAndDirectory("Workdir")
	if err != nil {
		return err
	}

	if commander.IsExistDir(filepath.Join(workdir, "von-network")) == false {
		err = commander.CloneFromURL(workdir, "https://github.com/Yukuro/von-network.git")
		if err != nil {
			return err
		}
		fmt.Printf("cloned VON-Network in %v\n", workdir)
	}

	//fmt.Printf("Your dot file is %v\n", dotFilepath)
	//fmt.Printf("Your workdir  is %v\n", workdir)

	//start docker-network
	fmt.Println("Building docker network...")
	fmt.Println("Enter network name")
	networkName, err := common.PromptString("NW name")
	if err != nil {
		return err
	}

	networkHash, err := commander.BuildDockerNetwork(networkName)
	if err != nil {
		return err
	}
	fmt.Printf("created %v : %v\n", networkName, networkHash)

	//Start VON-NW
	fmt.Println("[VON-Network]")
	err = ledger.RenameNetworks(workdir, networkName)
	fmt.Printf("Renamed networks: %v in %v/von-network/docker-compose.yml\n", networkName, workdir)

	fmt.Println("build and start network")
	err = commander.BuildVonNetwork(workdir)
	if err != nil {
		return err
	}
	err = commander.StartVonNetwork(workdir)
	if err != nil {
		return err
	}

	fmt.Println("Waiting for boot ledger...")
	time.Sleep(time.Second * 30)

	//register public DID to ledger
	fmt.Println("Registering to ledger...")
	agentNameList := analyse.GetAgentNameList(dotFilepath)

	agentNameAndSeed, err := commander.RegisterDID(agentNameList)
	//fmt.Println(agentSeedList)

	fmt.Println("[Agent]")
	//testtesttest := commander.IsExistDir(filepath.Join(workdir,"aries-cloudagent-python"))
	//fmt.Println(testtesttest)
	if commander.IsExistDir(filepath.Join(workdir, "aries-cloudagent-python")) == false {
		err = commander.CloneFromAcaPy(workdir, "https://github.com/Yukuro/aries-cloudagent-python.git")
		if err != nil {
			return err
		}
	}
	fmt.Println("Converting...")
	err = agent.ConvertFromGraph(dotFilepath, workdir, networkName, "192.168.3.15", agentNameAndSeed) //TODO get IP from command
	if err != nil {
		return err
	}

	return nil
}
