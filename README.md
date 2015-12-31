# About

Command line tool to validate JSON syntax of input files.

This tool simply exposes the super fast [encoding/json](https://golang.org/pkg/encoding/json/) to the command line.


# Installation

    go install github.com/martinlindhe/validjson


# Usage

    validjson file.json

    OK: file.json

To pretty print the result:

    validjson -p file.json

    {
        "bool": true
    }

# License

Under [MIT](LICENSE)
