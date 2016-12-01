package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Usage: odd_even [-h | --help] (ODD EVEN)...

Example, try:
  odd_even 1 2 3 4

Options:
  -h, --help`

	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
