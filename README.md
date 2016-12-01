docopt-go
=========

[![Build Status](https://travis-ci.org/aviddiviner/docopt-go.svg?branch=master)](https://travis-ci.org/aviddiviner/docopt-go)
[![Coverage Status](https://coveralls.io/repos/github/aviddiviner/docopt-go/badge.svg)](https://coveralls.io/github/aviddiviner/docopt-go)
[![GoDoc](https://godoc.org/github.com/aviddiviner/docopt-go?status.svg)](https://godoc.org/github.com/aviddiviner/docopt-go)

> Note: This is a work in progress refactor of the original package here: https://github.com/docopt/docopt.go

An implementation of [docopt](http://docopt.org/) in the [Go](http://golang.org/) programming language.

**docopt** helps you create *beautiful* command-line interfaces easily:

```go
package main

import (
	"fmt"
	"github.com/aviddiviner/docopt-go"
)

func main() {
	  usage := `Naval Fate.

Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

	  arguments, _ := docopt.ParseDoc(usage)
	  fmt.Println(arguments)
}
```

**docopt** parses command-line arguments based on a help message. Don't write parser code: a good help message already has all the necessary information in it.

## Installation

To use docopt in your Go code:

```go
import "github.com/aviddiviner/docopt-go"
```

To install docopt in your `$GOPATH`:

```console
$ go get github.com/aviddiviner/docopt-go
```

## Tests

All tests from the Python version are implemented and passing at [Travis CI](https://travis-ci.org/aviddiviner/docopt-go). New language-agnostic tests have been added to [test_golang.docopt](test_golang.docopt).

To run tests for docopt-go, use `go test`.

## API

Given a conventional command-line help message, docopt processes the arguments. See https://github.com/docopt/docopt#help-message-format for a description of the help message format.

This package exposes three different APIs, depending on the level of control required. The first, simplest way to parse your docopt usage is to just call:

```go
docopt.ParseDoc(usage)
```

This will use `os.Args[1:]` as the argv slice, and use the default parser options. If you want to provide your own version string and args, then use:

```go
docopt.ParseArgs(usage, argv, "1.2.3")
```

If the last parameter (version) is a non-empty string, it will be printed when `--version` is given in the argv slice. Finally, we can instantiate our own `docopt.Parser` which gives us control over how things like help messages are printed and whether to exit after displaying usage messages, etc.

```go
parser := &docopt.Parser{
  HelpHandler: docopt.PrintHelpOnly,
  OptionsFirst: true,
}
opts, err := parser.ParseArgs(usage, argv, "")
```

In particular, setting your own custom `HelpHandler` function makes unit testing your own docs with example command line invocations much more enjoyable.

All three of these return a map of option names to the values parsed from argv, and an error or nil. You can get the values using the helpers, or just treat it as a regular map:

```go
flag, _ := opts.Bool("--flag")
secs, _ := opts.Int("<seconds>")
```

More documentation is available at [godoc.org](https://godoc.org/github.com/aviddiviner/docopt-go).

## Unit Testing

Unit testing your own usage docs is recommended, so you can be sure that for a given command line invocation, the expected options are set. An example of how to do this is [in the examples folder](blob/master/examples/unit_test/unit_test.go).
