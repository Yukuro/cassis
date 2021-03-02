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
	"time"
)

// holderCmd represents the holder command
var holderCmd = &cobra.Command{
	Use:   "holder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doAsHolder()
	},
}

func init() {
	rootCmd.AddCommand(holderCmd)
}

func doAsHolder() error {
	menuList := []string{
		"build Agent",
		"accept invitation",
		"send proposal",
		"send request",
		"store credential",
	}

	currentAbsPath, err := filepath.Abs("./")
	if err != nil {
		return err
	}

	adminPort, err := commander.GetAdminPortFromWorkdir("holder", currentAbsPath)
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

	switch selectedMenu {
	case menuList[0]: // "build Agent"
		fmt.Println("build Agent...")
	case menuList[1]: // "accept invitation"
		invitation, err := common.PromptString("Enter invitation")
		if err != nil {
			return err
		}
		fmt.Printf("%v %v\n", targetUrl, invitation) //tmp
	case menuList[2]: // "send proposal"
		fmt.Println("send proposal...")
	case menuList[3]: // "send request"
		fmt.Println("send request...")
	case menuList[4]: // "store credential"
		fmt.Println("store credential...")

		credExIdList, err := commander.GetOfferCredExIdList(targetUrl)
		if err != nil {
			return err
		}

		//選択している時点でrequestする可否は問えているとみなす
		fmt.Printf("\n[offer box]\n")
		targetCredExId, err := common.PromptSelect("Select target cred_ex_id", credExIdList)
		if err != nil {
			return err
		}

		//request-credential
		err = commander.SendRequest(targetUrl, targetCredExId)
		if err != nil {
			return err
		}

		fmt.Println("Waiting for the procedure to next state...")
		time.Sleep(time.Second * 3)

		credId, err := commander.StoreCredential(targetUrl, targetCredExId)
		if err != nil {
			return err
		}

		fmt.Printf("\nCredential stored --> Credential ID :%v\n", credId)
	}

	return nil
}
