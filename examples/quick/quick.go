package main

import (
	"fmt"
	"github.com/ps78674/docopt.go"
)

func main() {
	usage := `Usage:
  quick tcp <host> <port> [--timeout=<seconds>]
  quick serial <port> [--baud=9600] [--timeout=<seconds>]
  quick -h | --help | --version`

	arguments, _ := docopt.ParseArgs(usage, nil, "0.1.1rc")
	fmt.Println(arguments)
}
