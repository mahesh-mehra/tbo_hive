package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"tbo_backend/objects"
)

func LoadConfig() {

	defer HandlePanic()

	jsonFile, err := os.Open("./local.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(byteValue, &objects.ConfigObj)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()

		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
}
