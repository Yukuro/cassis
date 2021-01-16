package test

import (
	"cli/src/analyse"
	"testing"
)

func TestBrowsing(t *testing.T){
	analyse.Browse("../test/test.dot")
}
