# About

Command line tool to validate JSON syntax of input files.

This tool simply exposes the super fast [encoding/json](https://golang.org/pkg/encoding/json/) to the command line.


# Installation

    go get -u github.com/martinlindhe/validjson


# Usage

    validjson file.json

    OK: file.json

To pretty print the result:

    validjson -p file.json

![screenshot](examples/pretty.png)


# License

Under [MIT](LICENSE)
