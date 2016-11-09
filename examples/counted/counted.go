package main

import (
	"fmt"
	"github.com/aviddiviner/docopt-go"
)

func main() {
	usage := `Usage: counted --help
       counted -v...
       counted go [go]
       counted (--path=<path>)...
       counted <file> <file>

Try: counted -vvvvvvvvvv
     counted go go
     counted --path ./here --path ./there
     counted this.txt that.txt`

	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
