package common

import (
	"errors"
	"github.com/manifoldco/promptui"
	"path/filepath"
)

func PromptForDot(name string) (string, error) {
	prompt := promptui.Prompt{
		Label: name,
		Validate: validateIsDotfile,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil{
		return "", err
	}
	return path, nil
}

func PromptForCsv(name string) (string, error) {
	prompt := promptui.Prompt{
		Label: name,
		Validate: validateIsCsvfile,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil{
		return "", err
	}
	return path, nil
}

func PromptForFileAndDirectory(name string) (string, error) {
	prompt := promptui.Prompt{
		Label: name,
		Validate: validateExistFileAndDirectory,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil{
		return "", err
	}
	return path, nil
}

func PromptConfirm(name string)(string, error){
	prompt := promptui.Prompt{
		Label: name,
		IsConfirm: true,
	}

	return prompt.Run()
}

func PromptString(name string)(string, error){
	prompt := promptui.Prompt{
		Label: name,
	}

	return prompt.Run()
}

func PromptNetworkName(name string)(string, error){
	prompt := promptui.Prompt{
		Label: name,
		Validate: validateExistDockerNetwork,
	}
	return prompt.Run()
}

func runAndConvAbs(p promptui.Prompt) (string, error){
	path, err := p.Run()
	if err != nil{
		return "", errors.New("can't enter dir")
	}
	path, err = filepath.Abs(path)
	if err != nil{
		return "", errors.New("can't convert relative to abs")
	}
	return path, nil
}
