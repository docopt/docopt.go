docopt.go
=========

Golang implementation of [docopt](http://docopt.org/) 0.6.1+fix

## Installation

import "github.com/kbatten/docopt.go" and then run `go get`

## API


```
args, output, err := docopt.Parse(doc string, argv []string, help bool, version string, optionsFirst bool)

- parse and return a map of args, output and all errors
```

```
args := docopt.ParseEasy(doc string)

- parse just doc and return a map of args
- handle all printing and non-fatal errors
- panic on fatal errors
- exit on user error or help
```

```
args, err := docopt.ParseQuiet(doc string, argv []string, help bool, version string, optionsFirst bool)

- parse and return a map of args and fatal errors
- handle printing of help
- exit on user error or help
```

```
args := docopt.ParseLoud(doc string, argv []string, help bool, version string, optionsFirst bool)

- parse and return a map of args
- handle all printing and non-fatal errors
- panic on fatal errors
- exit on user error or help
```

### arguments

`doc` usage string based on docopt language

`argv` optional argument vector. set to `nil` to use os.Args

`help` set to `true` to have docopt automatically handle `-h` and `--help`

`version` set to a non-empty string that will automatically be shown with `--version`

`optionsFirst` set to `true` to disallow mixing options and positional arguments

### return values

`args` map[string]interface{}. interface{} can be `bool`, `int`, `string`, `[]string`

`output` help output that would normally be displayed by the other `docopt.Parse*` functions.

`err`

- `nil`, no error
- `*docopt.UserError`, user argument error
- `*docopt.LanguageError`, developer error

## Example

``` go
package main

import (
        "fmt"
        docopt "github.com/kbatten/docopt.go"
)

func main() {
usage := `Naval Fate.

Usage:
  naval_fate.py ship new <name>...
  naval_fate.py ship <name> move <x> <y> [--speed=<kn>]
  naval_fate.py ship shoot <x> <y>
  naval_fate.py mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate.py -h | --help
  naval_fate.py --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

    arguments := docopt.ParseLoud(usage, nil, true, "Naval Fate 2.0", false)
    fmt.Println(arguments)
}
```

## Testing

All tests from the python version have been implemented pass

To run them use `go test`
