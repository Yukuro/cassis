package test

import (
	"cli/common"
	"cli/src/analyse"
	"fmt"
	"testing"
)

func TestBrowsing(t *testing.T) {
	analyse.Browse("../test/test.dot")
}

func TestConfirm(t *testing.T) {
	dotFilepath, err := common.PromptForDot("Dot")
	if err != nil {
		return
	}
	fmt.Println(dotFilepath)
}
