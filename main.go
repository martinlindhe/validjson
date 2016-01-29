package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/mattn/go-isatty"

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

		if isTerminal() {
			if b, err = highlight(b); err != nil {
				fmt.Println("ERROR highlight:", err)
			}
		}
		fmt.Printf(string(b))
		fmt.Println("")

	} else {
		fmt.Println("OK:", *inFile)
	}
}

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}

type rule struct {
	Expr    string
	Replace string
}

var rules = []rule{{"", ""},
	{`(?m)(\"[^\"]+\"):`, "\033[36m$1\033[39m:"},
	{`(?m)(^\s*[{}\]]{1}[,]*$|: [{\[])`, "\033[33m$1\033[39m"},
	{`: (\"[^\"]+\")`, ": \033[31m$1\033[39m$2"},
	{`(?m): ([\d][\d\.e+]*)([,]*)$`, ": \033[33m$1\033[39m"},
	{`(?m)(^\s+(?:[\d][\d\.e+]*|true|false|null))([,]*)$`, "\033[35m$1\033[39m$2"},
	{`(?::) ((?:true|false|null))`, ": \033[31m$1\033[39m"},
	{`(?m)^(true|false|null)$`, "\033[31m$1\033[39m"},
}

func highlight(data []byte) ([]byte, error) {
	var (
		re  *regexp.Regexp
		err error
	)
	for _, rule := range rules {
		if rule.Expr == "" {
			continue
		}

		re, err = regexp.Compile(rule.Expr)
		if err != nil {
			return nil, err
		}
		data = re.ReplaceAll(data, []byte(rule.Replace))
	}

	return data, nil
}
