package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file.json>\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	inFile := os.Args[1]

	if !fileExists(inFile) {
		fmt.Printf("Error: file %s not found.\n", inFile)
		os.Exit(1)
	}

	data, _ := ioutil.ReadFile(inFile)

	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf("JSON error in %s: %s\n", inFile, err)
		os.Exit(1)
	}

	fmt.Printf("OK: %s\n", inFile)
}

func fileExists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
