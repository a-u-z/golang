package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// fileName aa.json
func readJson(fileName string) ([]byte, error) {
	// 開啟json檔案
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
