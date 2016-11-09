package main

import (
	"fmt"
	"github.com/aviddiviner/docopt-go"
)

func main() {
	usage := `Not a serious example.

Usage:
  calculator <value> ( ( + | - | * | / ) <value> )...
  calculator <function> <value> [( , <value> )]...
  calculator (-h | --help)

Examples:
  calculator 1 + 2 + 3 + 4 + 5
  calculator 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'
  calculator sum 10 , 20 , 30 , 40

Options:
  -h, --help
`
	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
