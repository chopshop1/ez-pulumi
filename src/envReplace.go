package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func validate(input string) error {
	stringLength := len(input)
	if stringLength == 0 {
		return errors.New("Please Input a file path")
	}
	return nil
}

func getFiles() []string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	var fileArr []string
	for _, file := range files {
		fileArr = append(fileArr, file.Name())
	}
	return fileArr
}

func EnvReplace() {
	prompt := promptui.Select{
		Label: "Which env file would you like to use?",
		Items: getFiles(),
	}
	_, result, errp := prompt.Run()

	if errp != nil {
		fmt.Printf("Prompt failed %v\n", errp)
	}
	file := result
	ReplaceEnv(file)
}

func ReplaceEnv(file string) error {
	envVars, _ := ioutil.ReadFile(file)
	var data map[string]map[string]interface{}

	err := json.Unmarshal(envVars, &data)
	if err != nil {
		log.Fatal("error parsing json")
		return err
	}

	for envVar := range data {
		envVarValue := fmt.Sprint(data[envVar]["value"])
		isSecret := data[envVar]["secret"]
		pulumiCmd := []string{"config", "set", fmt.Sprint(envVar), fmt.Sprint(envVarValue)}

		if isSecret == true {
			//Inserts --secret flag into array
			pulumiCmd = append(pulumiCmd[:2], append([]string{"--secret"}, pulumiCmd[2:]...)...)
		}

		log.Print(pulumiCmd)

		_, err := exec.Command("pulumi", pulumiCmd...).Output()
		if err != nil {
			log.Fatal("error:", err)
			return err
		}
	}
	return nil
}
