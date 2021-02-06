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
package commander

import (
	"cli/common"
)

func RemoveSystem() error {
	workdir, err := common.PromptForFileAndDirectory("Enter target workdir")
	if err != nil {
		return err
	}

	err = DockerComposeDownAtWorkdir(workdir)
	if err != nil {
		return err
	}

	err = DockerComposeDownAtVonNw(workdir)
	if err != nil {
		return err
	}

	//rmConf, err := common.PromptConfirm("Remove all docker images?")
	//if err != nil{
	//	return err
	//}
	//if rmConf == "y"{
	//	err = commander.RemoveAllDockerImages()
	//	if err != nil{
	//		return err
	//	}
	//}

	//err = commander.RemoveAllFilesAtWorkDir(workdir)
	//if err != nil{
	//	return err
	//}

	return nil
}
