package main

import (
	"testing"
)

func TestEnv(t *testing.T) {
	output := ReplaceEnv("testFiles/test.json")
	if output != nil {
		t.Error("EnvReplace had logged an error")
	}
}
