package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inFile = kingpin.Arg("file", "JSON file").Required().ExistingFile()
	pretty = kingpin.Flag("pretty", "Pretty print result").Short('p').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, _ := ioutil.ReadFile(*inFile)

	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		fmt.Println("ERROR:", *inFile, err)
		os.Exit(1)
	}

	if *pretty {
		b, err := json.MarshalIndent(f, "", "    ")
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Printf(string(b))

	} else {
		fmt.Println("OK:", *inFile)
	}
}
