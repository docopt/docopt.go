package docopt

import (
	"fmt"
	"sort"
)

func ExampleParseArgs() {
	usage := `Usage:
  example tcp [<host>] [--force] [--timeout=<seconds>]
  example serial <port> [--baud=<rate>] [--timeout=<seconds>]
  example --help | --version`

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
	//    <host> 127.0.0.1
	//    <port> <nil>
	//    serial false
	//       tcp true
}

func ExampleOpts_Bind() {
	usage := `Usage:
  example tcp [<host>] [--force] [--timeout=<seconds>]
  example serial <port> [--baud=<rate>] [--timeout=<seconds>]
  example --help | --version`

	type options struct {
		Tcp     bool   `docopt:"tcp"`
		Serial  bool   `docopt:"serial"`
		Host    string `docopt:"<host>"`
		Port    int    `docopt:"<port>"`
		Force   bool
		Timeout int
		Baud    int
	}

	// Parse the command line `example tcp 127.0.0.1 --force`
	argv := []string{"tcp", "127.0.0.1", "--force"}
	arguments, _ := ParseArgs(usage, argv, "0.1.1rc")

	var opts options
	arguments.Bind(&opts)
	fmt.Printf("%+v\n", opts)

	// Parse the command line `example serial 443 --baud=9600`
	argv = []string{"serial", "443", "--baud=9600"}
	arguments, _ = ParseArgs(usage, argv, "0.1.1rc")

	opts = options{}
	arguments.Bind(&opts)
	fmt.Printf("%+v\n", opts)

	// Output:
	// {Tcp:true Serial:false Host:127.0.0.1 Port:0 Force:true Timeout:0 Baud:0}
	// {Tcp:false Serial:true Host: Port:443 Force:false Timeout:0 Baud:9600}
}
