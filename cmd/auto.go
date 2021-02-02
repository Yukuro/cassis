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

// autoBuildCmd represents the build command
var autoBuildCmd = &cobra.Command{
	Use:   "auto",
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
	rootCmd.AddCommand(autoBuildCmd)
}

func buildAgent() error {
	workdir, err := common.PromptForFileAndDirectory("Enter your workdir")
	if err != nil {
		return err
	}

	//get .dot file
	//fmt.Println("Enter your dot file...")
	dotFilepath, err := common.PromptForDot("Enter your DOT file dir")
	if err != nil {
		return err
	}

	printProgressWithDot("Analysing", 3)
	err = analyse.Browse(dotFilepath)
	if err != nil {
		return err
	}

	//Does Issuer - Holder =~ VC exist ?
	vcRequired, err := analyse.RequireVC(dotFilepath)
	if err != nil {
		return err
	}
	if vcRequired {
		//fmt.Println("detected Issuer - Holder relation...")
		printWithNewLine("detected Issuer - Holder relation..")

		//csvFilepath, err := common.PromptForCsv("Enter your CSV file dir")
		_, err := common.PromptForCsv("Enter your CSV file dir") //TODO CSVの処理書く
		if err != nil {
			return err
		}
		//fmt.Printf("Your csv file is %v\n", csvFilepath)
	}

	printWithNewLine("Start building Network and Agent...")

	if commander.IsExistDir(filepath.Join(workdir, "von-network")) == false {
		err = commander.CloneFromURL(workdir, "https://github.com/Yukuro/von-network.git")
		if err != nil {
			return err
		}
		//fmt.Printf("cloned VON-Network in %v\n", workdir)
	}

	//fmt.Printf("Your dot file is %v\n", dotFilepath)
	//fmt.Printf("Your workdir  is %v\n", workdir)

	//start docker-network
	//fmt.Println("Building docker network...")
	//fmt.Println("Enter network name")
	networkName, err := common.PromptString("Enter the docker network name")
	if err != nil {
		return err
	}

	networkHash, err := commander.BuildDockerNetwork(networkName)
	if err != nil {
		return err
	}
	fmt.Printf("created %v : %v\n", networkName, networkHash)

	//Start VON-NW
	fmt.Println("[Ledger]")
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

	//fmt.Println("Waiting for boot ledger...")
	printWithNewLine("Wait 30 seconds for the ledger to start.")
	time.Sleep(time.Second * 30)

	//register public DID to ledger
	//fmt.Println("Registering to ledger...")
	agentNameList := analyse.GetAgentNameList(dotFilepath)

	agentNameAndSeed, err := commander.RegisterDID(agentNameList)
	//fmt.Println(agentSeedList)

	//fmt.Println("[Agent]")
	printWithNewLine("[Agent]")
	//testtesttest := commander.IsExistDir(filepath.Join(workdir,"aries-cloudagent-python"))
	//fmt.Println(testtesttest)
	if commander.IsExistDir(filepath.Join(workdir, "aries-cloudagent-python")) == false {
		err = commander.CloneFromAcaPy(workdir, "https://github.com/Yukuro/aries-cloudagent-python.git")
		if err != nil {
			return err
		}
	}
	err = agent.ConvertFromGraph(dotFilepath, workdir, networkName, "192.168.3.15", agentNameAndSeed) //TODO get IP from command
	if err != nil {
		return err
	}
	printWithNewLine("Rewrote docker-compose.yml")

	fmt.Println("Up Agent...")
	err = commander.DockerComposeUpAtWorkdir(workdir)
	if err != nil {
		panic(err)
	}

	fmt.Println("done!")
	return nil
}

func printWithNewLine(comment string) {
	fmt.Printf("\n%v\n", comment)
}

func printProgressWithDot(comment string, times int) {
	fmt.Printf("\n%v", comment)
	for i := 0; i < times; i++ {
		fmt.Printf(".")
		time.Sleep(time.Second * 1)
	}
	fmt.Println()
}
