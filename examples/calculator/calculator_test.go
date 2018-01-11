package main

import (
	"github.com/docopt/docopt-go/examples"
)

func Example() {
	examples.TestUsage(usage, "calculator 1 + 2 + 3 + 4 + 5")
	examples.TestUsage(usage, "calculator 1 + 2 * 3 / 4 - 5")
	examples.TestUsage(usage, "calculator sum 10 , 20 , 30 , 40")
	// Output:
	//         * 0
	//         + 4
	//         , 0
	//         - 0
	//    --help false
	//         / 0
	// <function> <nil>
	//   <value> [1 2 3 4 5]
	//
	//         * 1
	//         + 1
	//         , 0
	//         - 1
	//    --help false
	//         / 1
	// <function> <nil>
	//   <value> [1 2 3 4 5]
	//
	//         * 0
	//         + 0
	//         , 3
	//         - 0
	//    --help false
	//         / 0
	// <function> sum
	//   <value> [10 20 30 40]
}
