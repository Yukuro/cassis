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
	"cli/src/agent"
	"cli/src/commander"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return initEnv()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initEnv() error {
	if commander.IsExistDir(".cassis") {
		fmt.Println(".cassis already exists")
		fmt.Println("Is that working directory correct?")
		return nil
	}

	//defaultUmask := syscall.Umask(0)
	if err := mkdirInitDir(); err != nil {
		return err
	}
	//syscall.Umask(defaultUmask)

	if err := agent.CreateNilConf(".cassis"); err != nil {
		return err
	}
	return nil
}

func mkdirInitDir() error {
	// TODO しかるべき権限に設定する
	if err := os.Mkdir(".cassis", 0777); err != nil {
		return err
	}
	return nil
}
