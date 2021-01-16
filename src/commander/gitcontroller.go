package commander

import (
	"errors"
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
	"strings"
)

//TODO override "option"
func CloneFromURL(workdir string, url string) error {
	parameter := []string{
		"clone",
		url,
	}

	os.Chdir(workdir)

	cmd := exec.Command("git", parameter...)
	fl, err := pty.Start(cmd)
	if err != nil{
		errors.New("can't clone from url")
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func CloneFromAcaPy(workdir string, url string) error {
	parameter := []string{
		"clone",
		url,
		"-b",
		"0.5.6",
	}

	os.Chdir(workdir)

	cmd := exec.Command("git", parameter...)
	fl, err := pty.Start(cmd)
	if err != nil{
		errors.New("can't clone from url")
	}
	io.Copy(os.Stdout, fl)
	return nil
}

func ExtractFolderName(url string, owner string) string {
	forward := "https://github.com/"
	back := ".git"

	removedForward := strings.Replace(url, forward + owner + "/", "", -1)
	result := strings.Replace(removedForward, back, "", -1)

	return result
}
