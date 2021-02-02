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
	"fmt"
	"github.com/spf13/cobra"
)

// issuerCmd represents the issuer command
var issuerCmd = &cobra.Command{
	Use:   "issuer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doAsIssuer()
	},
}

func init() {
	rootCmd.AddCommand(issuerCmd)
}

func doAsIssuer() error {
	menuList := []string{
		"build Agent",
		"register schema",
		"register credential definition",
	}

	//DOT言語が同ディレクトリに存在するか？


	selectedMenu, err := common.PromptSelect("Select", menuList)
	if err != nil{
		return err
	}
	switch selectedMenu {
	case menuList[0]: // "build Agent"
		fmt.Println("build Agent...")
	case menuList[1]: // "register schema"
		fmt.Println("register schema...")
	case menuList[2]: // "register credential definition"
		fmt.Println("register credential definition")
	}
	return nil
}
