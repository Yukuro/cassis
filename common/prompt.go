package common

import (
	"errors"
	"github.com/manifoldco/promptui"
	"path/filepath"
)

func PromptForDot(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateIsDotfile,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil {
		return "", err
	}
	return path, nil
}

func PromptForCsv(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateIsCsvfile,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil {
		return "", err
	}
	return path, nil
}

func PromptForFileAndDirectory(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateExistFileAndDirectory,
	}

	path, err := runAndConvAbs(prompt)
	if err != nil {
		return "", err
	}
	return path, nil
}

func PromptConfirm(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:     name,
		IsConfirm: true,
	}

	return prompt.Run()
}

func PromptSelect(name string, item []string) (string, error) {
	prompt := promptui.Select{
		Label: name,
		Items: item,
	}

	_, result, err := prompt.Run()
	return result, err
}

func PromptString(name string) (string, error) {
	prompt := promptui.Prompt{
		Label: name,
	}

	return prompt.Run()
}

func PromptNetworkName(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateExistDockerNetwork,
	}
	return prompt.Run()
}

func PromptYesOrNo(name string) (bool, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateYesOrNo,
	}
	res, err := prompt.Run()
	if err != nil {
		return false, err
	}

	if res == "Y" || res == "y" {
		return true, nil
	}
	return false, nil
}

func runAndConvAbs(p promptui.Prompt) (string, error) {
	path, err := p.Run()
	if err != nil {
		return "", errors.New("can't enter dir")
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return "", errors.New("can't convert relative to abs")
	}
	return path, nil
}
