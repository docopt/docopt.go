package main

import (
	"github.com/docopt/docopt-go/examples"
)

func Example() {
	examples.TestUsage(usage, "counted -vvvvvvvvvv")
	examples.TestUsage(usage, "counted go go")
	examples.TestUsage(usage, "counted --path ./here --path ./there")
	examples.TestUsage(usage, "counted this.txt that.txt")
	// Output:
	//    --help false
	//    --path []
	//        -v 10
	//    <file> []
	//        go 0
	//
	//    --help false
	//    --path []
	//        -v 0
	//    <file> []
	//        go 2
	//
	//    --help false
	//    --path [./here ./there]
	//        -v 0
	//    <file> []
	//        go 0
	//
	//    --help false
	//    --path []
	//        -v 0
	//    <file> [this.txt that.txt]
	//        go 0
}
