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
	"cli/src/agent"
	"cli/src/commander"
	"fmt"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
		"invite holder",
		"schema",
		"credential definition",
	}

	currentAbsPath, err := filepath.Abs("./")
	if err != nil {
		return err
	}

	//DOT言語が同ディレクトリに存在するか？
	//auto modeでない限り、そんなに意味ないような...ファイル内容の同期が難しそう
	isExistDotInCurrentDir, err := commander.IsExistDotInDir(currentAbsPath)
	if err != nil {
		return err
	}
	if isExistDotInCurrentDir {
		fmt.Println("Dot file exists")
	}

	//Issuer選択
	//network inviteから流用
	//将来的には.cassis/config.ymlから読む
	adminPort, err := commander.GetAdminPortFromWorkdir(currentAbsPath)
	if err != nil {
		return err
	}

	selectedAgent, err := common.PromptSelect("Select target agent", commander.MapKeyToSlice(adminPort))
	if err != nil {
		return err
	}
	//fmt.Println(selectedAgent)
	targetPort := adminPort[selectedAgent]
	targetUrl := "http://localhost:" + targetPort

	selectedMenu, err := common.PromptSelect("Select", menuList)
	if err != nil {
		return err
	}
	switch selectedMenu {
	case menuList[0]: // "build Agent"
		fmt.Println("build Agent...")
	case menuList[1]: // "invite Holder"
		fmt.Println("invite Holder...")
		connectionId, invitation, err := commander.InvitationToHolder(targetUrl, selectedAgent)
		if err != nil {
			return err
		}

		fmt.Printf("\nConnection ID: %v\n", connectionId)
		fmt.Println("Scan the following qr code with your app")
		//fmt.Println(invitation)

		fmt.Printf("\n%v\n", invitation)
		qrterminal.GenerateHalfBlock(invitation, qrterminal.L, os.Stdout)
	case menuList[2]: // "register schema"
		fmt.Println("originate schema...")
		schemaMenuList := []string{
			"configure schema",
			"originate schema",
		}
		selectedSchemaMenu, err := common.PromptSelect("Select", schemaMenuList)
		if err != nil {
			return err
		}
		switch selectedSchemaMenu {
		case schemaMenuList[0]: // "configure schema"
			fmt.Println("set attribute")
			schemaName, err := common.PromptString("schema name")
			if err != nil {
				return err
			}
			var schemaAttribute []string
			fmt.Println("Type \"end\" to exit.")
			for {
				attr, err := common.PromptString("attribute")
				if err != nil {
					return err
				}

				if attr == "end" {
					break
				}

				schemaAttribute = append(schemaAttribute, attr)
			}

			// 一旦Attributeをymlに保存
			//sc := map[string]string{
			//	"name": schemaName,
			//	"version": "1.0",
			//	"attr": schemaAttribute,
			//}
			err = agent.CreateIssuerConf(".cassis", schemaName, "1.0", schemaAttribute, "", "")
			if err != nil {
				return err
			}

			// originateを望むならoriginateしてschemaIdを取得
			//fmt.Printf("\nDo you want to originate now?\n")
			isOriginate, err := common.PromptYesOrNo("Do you want to originate now?")
			if !isOriginate {
				return nil
			}

			// schemaNameとschemaAttributeをconfig.ymlに書き込む
			//fmt.Printf("%v %v\n", schemaName, schemaAttribute)
			originatedSchemaName, originatedSchemaVersion, originatedSchemaAttr, originatedSchemaId, err :=
				commander.OriginateSchema(targetUrl+"/schemas", schemaName, "1.0", schemaAttribute)
			if err != nil {
				return err
			}

			err = agent.CreateIssuerConf(".cassis", originatedSchemaName, originatedSchemaVersion, originatedSchemaAttr, originatedSchemaId, "")
			if err != nil {
				return err
			}

			//fmt.Printf("targetURL %v", targetUrl)
			//fmt.Println("done!")
			fmt.Printf("\nschema: %v --> ledger originated!\n", schemaName)

		case schemaMenuList[1]: // "originate schema"
			fmt.Println("originate schema")

			conf, err := agent.AnalyzeIssuerConf(".cassis")
			if err != nil {
				return err
			}

			originateSchemaMenuList := []string{}
			for sN, sA := range conf {
				tp := fmt.Sprintf("%v %v", sN, sA)
				originateSchemaMenuList = append(originateSchemaMenuList, tp)
			}

			selectedOriginateSchema, err := common.PromptSelect("Select schema", originateSchemaMenuList)
			fmt.Printf("%v\n", selectedOriginateSchema)

			//選択したschemaのパーサー書く
			//選択したschemaでoriginate
		}
	case menuList[3]: // "register credential definition"
		fmt.Println("originate credential definition")
		conf, err := agent.AnalyzeIssuerConf(".cassis")
		if err != nil {
			return err
		}

		var cred_defMenuList []string
		for _, sc := range conf {
			menu := fmt.Sprintf("%v(%v) %v", sc["name"], sc["version"], sc["id"])
			cred_defMenuList = append(cred_defMenuList, menu)
		}

		selected, err := common.PromptSelect("Select", cred_defMenuList)
		if err != nil {
			return nil
		}

		var selectedIndex int
		for i, mn := range cred_defMenuList {
			if mn == selected {
				selectedIndex = i
				break
			}
		}

		sc := conf[selectedIndex]
		//fmt.Printf("%v %v %v\n", sc["name"], sc["version"], sc["id"])

		originatedCred_defId, err := commander.OriginateCred_def(targetUrl+"/credential-definitions", sc["id"])
		if err != nil {
			return err
		}

		err = agent.AddIssuerConfWithWorkdir(".cassis", "", "", []string{}, "", originatedCred_defId)
		if err != nil {
			return err
		}

		fmt.Printf("credential definition ( %v ) --> ledger done!", originatedCred_defId)
	}
	return nil
}
