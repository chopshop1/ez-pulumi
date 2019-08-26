package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Which command would you like to use?",
		Items: []string{"Move Env", "Clear Pending Operations"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Move Env":
		EnvReplace()
	case "Clear Pending Operations":
		ClearPendingOperations()
	}
}
