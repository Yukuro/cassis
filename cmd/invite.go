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
*/
package cmd

import (
	"cli/common"
	"cli/src/commander"
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return inviteHolder()
	},
}

func init() {
	networkCmd.AddCommand(inviteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inviteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inviteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func inviteHolder() error {
	workdir, err := common.PromptForFileAndDirectory("Enter workdir")
	if err != nil {
		return err
	}

	adminPort, err := commander.GetAdminPortFromWorkdir(workdir)
	if err != nil {
		return err
	}

	selectedAgent, err := common.PromptSelect("Select target agent", commander.MapKeyToSlice(adminPort))
	if err != nil {
		return err
	}
	//fmt.Println(selectedAgent)

	targetPort := adminPort[selectedAgent]
	inviteUrl := "http://localhost:" + targetPort

	//fmt.Printf("POST to %v\n", inviteUrl)

	connectionId, invitation, err := commander.InvitationToHolder(inviteUrl, selectedAgent)
	if err != nil {
		return err
	}

	fmt.Printf("\nConnection ID: %v\n", connectionId)
	fmt.Println("Scan the following qr code with your app")
	//fmt.Println(invitation)

	qrterminal.GenerateHalfBlock(invitation, qrterminal.L, os.Stdout)

	return nil
}
