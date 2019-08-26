package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	file := os.Args[1]
	EnvReplace(file)
}

func EnvReplace(file string) error {
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

		cmd := exec.Command("pulumi", pulumiCmd...)

		_, err := cmd.Output()
		if err != nil {
			log.Fatal("error:", err)
			return err
		}
	}
	return nil
}
