package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	output := EnvReplace("test.json")
	if output != nil {
		t.Error("EnvReplace had logged an error")
	}
}
