docopt.go
=========

Golang implementation of [docopt](http://docopt.org/) 0.6.1+fix

## Installation

import "github.com/docopt/docopt.go" and then run `go get`.

## API

``` go
func docopt.Parse(doc string, argv []string, help bool, version string, optionsFirst bool)
(args map[string]interface{}, err error)
```

Parse `argv` based on command-line interface described in `doc`.

docopt creates your command-line interface based on its description that you pass as `doc`. Such description can contain --options, <positional-argument>, commands, which could be [optional], (required), (mutually | exclusive) or repeated...

### arguments

`doc` Description of your command-line interface.

`argv` Argument vector to be parsed. os.Args[1:] is used if nil.

`help` Set to false to disable automatic help on -h or --help options..

`version` If set to something besides an empty string, the string will be printed
 if --version is in argv.

`optionsFirst` Set to true to require options precede positional arguments,
 i.e. to forbid options and positional arguments intermix..

### return values

`args`, map[string]interface{}. A map, where keys are names of command-line elements such as e.g. "--verbose" and "<path>", and values are the  parsed values of those elements. interface{} can be `bool`, `int`, `string`, `[]string`.

`err`, error. Either *docopt.LanguageError, *docopt.UserError or nil

## Example

``` go
package main

import (
    "fmt"
    "github.com/docopt/docopt.go"
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

    arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
    fmt.Println(arguments)
}
```

## Testing

All tests from the python version have been implemented and all are passing.

New language agnostic tests have been added to `test_golang.docopt`.

To run them use `go test`.
