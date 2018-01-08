# About

Command line tool to validate and pretty-print JSON syntax of
input files.


## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/validjson/releases)

Or install from source:

    go get -u github.com/martinlindhe/validjson


## Usage

Exit code will be 0 if file is good.

    $ validjson file.json
    OK: file.json

    $ curl http://site.com/file.json | validjson
    OK: -


## Pretty-print

    $ validjson -p file.json

![screenshot](examples/pretty.png)


## License

Under [MIT](LICENSE)
