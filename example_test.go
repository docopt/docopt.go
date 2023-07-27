package docopt

import (
	"fmt"
	"net"
	"sort"
)

func ExampleParseArgs() {
	usage := `Usage:
  example tcp [<host>...] [--force] [--timeout=<seconds>]
  example serial <port> [--baud=<rate>] [--timeout=<seconds>]
  example --help | --version`

	// Parse the command line `example tcp 127.0.0.1 --force`
	argv := []string{"tcp", "127.0.0.1", "--force"}
	opts, _ := ParseArgs(usage, argv, "0.1.1rc")

	// Sort the keys of the options map
	var keys []string
	for k := range opts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print the option keys and values
	for _, k := range keys {
		fmt.Printf("%9s %v\n", k, opts[k])
	}

	// Output:
	//    --baud <nil>
	//   --force true
	//    --help false
	// --timeout <nil>
	// --version false
	//    <host> [127.0.0.1]
	//    <port> <nil>
	//    serial false
	//       tcp true
}

func ExampleOpts_Bind() {
	usage := `docopt-go example

Usage:
  example tcp [<host>...] [--force] [--timeout=<seconds>] [--gateway=<ip>]
  example serial <port> [--baud=<rate>] [--timeout=<seconds>]
  example --help | --version

Options:
  --gateway <ip>  Set the gateway IP address [default: 192.168.0.1]
  -h --help       Show this screen.
`

	// Parse the command line `example serial 443 --baud=9600`
	argv := []string{"serial", "443", "--baud=9600"}
	opts, _ := ParseArgs(usage, argv, "0.1.1rc")

	var conf struct {
		Tcp     bool
		Serial  bool
		Host    []string
		Port    int
		Force   bool
		Timeout int
		Baud    int
		Gateway net.IP
	}
	opts.Bind(&conf)

	if conf.Serial {
		fmt.Printf(
			"port: %d, baud: %d, gateway: %v",
			conf.Port,
			conf.Baud,
			conf.Gateway,
		)
	}

	// Output:
	// port: 443, baud: 9600, gateway: 192.168.0.1
}
