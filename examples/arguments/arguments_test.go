package main

import (
	"github.com/docopt/docopt-go/examples"
)

func Example() {
	examples.TestUsage(usage, "arguments -qv")
	examples.TestUsage(usage, "arguments --left file.A file.B")
	// Output:
	//    --help false
	//    --left false
	//   --right false
	//        -q true
	//        -r false
	//        -v true
	// CORRECTION <nil>
	//      FILE []
	//
	//    --help false
	//    --left true
	//   --right false
	//        -q false
	//        -r false
	//        -v false
	// CORRECTION file.A
	//      FILE [file.B]
}
