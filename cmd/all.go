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
	"cli/src/agent-build"
	"cli/src/commander"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doAsAll()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

func doAsAll() error {
	menuList := []string{
		"build Agent",
		"remove Agent",
	}

	selectedMenu, err := common.PromptSelect("Select", menuList)
	if err != nil {
		return err
	}
	switch selectedMenu {
	case menuList[0]: // "build Agent
		if err := agent_build.BuildAgent(); err != nil {
			return err
		}
	case menuList[1]: // "remove Agent"
		if err := commander.RemoveSystem(); err != nil {
			return err
		}
	}
	return nil
}
