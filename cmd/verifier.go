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
	"github.com/spf13/cobra"
	"path/filepath"
)

// verifierCmd represents the verifier command
var verifierCmd = &cobra.Command{
	Use:   "verifier",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doAsVerifier()
	},
}

func init() {
	rootCmd.AddCommand(verifierCmd)
}

func doAsVerifier() error {
	menuList := []string{
		"request presentation",
		"verify presentation",
	}

	currentAbsPath, err := filepath.Abs("./")
	if err != nil {
		return err
	}

	adminPort, err := commander.GetAdminPortFromWorkdir("verifier", currentAbsPath)
	if err != nil {
		return err
	}

	selectedAgent, err := common.PromptSelect("Select target agent", commander.MapKeyToSlice(adminPort))
	if err != nil {
		return err
	}
	targetPort := adminPort[selectedAgent]
	targetUrl := "http://localhost:" + targetPort

	selectedMenu, err := common.PromptSelect("Select", menuList)
	if err != nil {
		return err
	}

	//TODO remove
	fmt.Printf("Your target URL is %v\n", targetUrl)

	switch selectedMenu {
	case "request presentation":
		fmt.Printf("request-presentaion\n")

	case "verify presentation":
		fmt.Printf("verify presentation\n")
	}

	return nil
}
