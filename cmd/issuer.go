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
		"schema",
		"credential definition",
	}

	//DOT言語が同ディレクトリに存在するか？
	//auto modeでない限り、そんなに意味ないような...ファイル内容の同期が難しそう
	isExistDotInCurrentDir, err := commander.IsExistDotInDir(".")
	if err != nil{
		return err
	}
	if isExistDotInCurrentDir {
		fmt.Println("Dot file exists")
	}

	selectedMenu, err := common.PromptSelect("Select", menuList)
	if err != nil{
		return err
	}
	switch selectedMenu {
	case menuList[0]: // "build Agent"
		fmt.Println("build Agent...")
	case menuList[1]: // "register schema"
		fmt.Println("originate schema...")
		schemaMenuList := []string{
			"configure schema",
			"originate schema",
		}
		selectedSchemaMenu, err := common.PromptSelect("Select", schemaMenuList)
		if err != nil{
			return err
		}
		switch selectedSchemaMenu {
		case schemaMenuList[0]: // "configure schema"
			fmt.Println("set attribute")
			schemaName, err := common.PromptString("Set schema name")
			if err != nil{
				return err
			}
			var schemaAttribute []string
			fmt.Println("Type \"end\" to exit.")
			for {
				attr, err := common.PromptString("Set attribute")
				if err != nil{
					return err
				}

				if attr == "end" {
					break
				}

				schemaAttribute = append(schemaAttribute, attr)
			}

			// schemaNameとschemaAttributeをconfig.ymlに書き込む
			fmt.Printf("%v %v\n", schemaName, schemaAttribute)

			// originateするか聞く
		case schemaMenuList[1]: // "originate schema"
			fmt.Println("originate schema")
		}
	case menuList[2]: // "register credential definition"
		fmt.Println("originate credential definition")
	}
	return nil
}
