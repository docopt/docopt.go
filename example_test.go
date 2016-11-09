package docopt

import (
	"fmt"
	"sort"
)

func ExampleParseArgs() {
	usage := `Usage:
  example tcp [<host>] [--force] [--timeout=<seconds>]
  example serial <port> [--baud=<rate>] [--timeout=<seconds>]
  example -h | --help | --version`

	// Parse the command line `example tcp 127.0.0.1 --force`
	argv := []string{"tcp", "127.0.0.1", "--force"}
	arguments, _ := ParseArgs(usage, argv, "0.1.1rc")

	// Sort the keys of the arguments map
	var keys []string
	for k := range arguments {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print the argument keys and values
	for _, k := range keys {
		fmt.Printf("%9s %v\n", k, arguments[k])
	}

	// Output:
	//    --baud <nil>
	//   --force true
	//    --help false
	// --timeout <nil>
	// --version false
	//        -h false
	//    <host> 127.0.0.1
	//    <port> <nil>
	//    serial false
	//       tcp true
}
