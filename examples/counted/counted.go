package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

var usage = `Usage: counted --help
       counted -v...
       counted go [go]
       counted (--path=<path>)...
       counted <file> <file>

Try: counted -vvvvvvvvvv
     counted go go
     counted --path ./here --path ./there
     counted this.txt that.txt`

func main() {
	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
