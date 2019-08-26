package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type PulumiStackInterface struct {
	Manifest          interface{}   `json:"manifest"`
	SecretsProviders  interface{}   `json:"secrets_providers"`
	Resources         []interface{} `json:"resources"`
	PendingOperations []string      `json:"pending_operations"`
}
type PulumiStack struct {
	Version    int                  `json:"version"`
	Deployment PulumiStackInterface `json:"deployment"`
}

func WriteDataToFileAsJSON(data interface{}, filedir string) (int, error) {
	//write data as buffer to json encoder
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return 0, err
	}
	file, err := os.OpenFile(filedir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return 0, err
	}
	n, err := file.Write(buffer.Bytes())
	if err != nil {
		return 0, err
	}
	return n, nil
}

func HandelError(err error) error {
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func ClearPendingOperations() error {
	cmdOutput, errz := exec.Command("pulumi", "stack", "export").Output()
	HandelError(errz)

	data := new(PulumiStack)

	err := json.Unmarshal(cmdOutput, data)
	HandelError(err)

	if len(data.Deployment.PendingOperations) != 0 {
		data.Deployment.PendingOperations = []string{}
		_, err = WriteDataToFileAsJSON(data, "stack.json")
		HandelError(err)
		//TODO:see if --force is necessary
		stackImportCmdOutput, stackImportCmdErr := exec.Command("pulumi", "stack", "import", "--file", "./stack.json", "--force").Output()
		HandelError(stackImportCmdErr)
		fmt.Println(string(stackImportCmdOutput))

		removeStackFileErr := exec.Command("rm", "./stack.json").Run()
		HandelError(removeStackFileErr)
	} else {
		fmt.Println("There are no pennding operations.")
	}

	return nil
}
