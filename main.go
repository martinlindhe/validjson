package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/alecthomas/kingpin/v2"
	termutil "github.com/andrew-d/go-termutil"
	"github.com/mattn/go-isatty"
)

var (
	inFile = kingpin.Arg("file", "JSON file.").String()
	pretty = kingpin.Flag("pretty", "Pretty print result.").Short('p').Bool()
	quiet  = kingpin.Flag("quiet", "Don't output on success.").Short('q').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, err := readPipeOrFile(*inFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	filename := "-"
	if *inFile != "" {
		filename = *inFile
	}

	var f interface{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		fmt.Println("ERROR:", filename, err)
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
		if !*quiet {
			fmt.Println("OK:", filename)
		}
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

// readPipeOrFile reads from stdin if pipe exists, else from provided file
func readPipeOrFile(fileName string) ([]byte, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		return ioutil.ReadAll(os.Stdin)
	}
	if fileName == "" {
		return nil, fmt.Errorf("no piped data and no file provided")
	}
	return ioutil.ReadFile(fileName)
}
