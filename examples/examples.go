package examples

import (
	"fmt"
	"sort"
	"strings"

	"github.com/docopt/docopt-go"
)

// TestUsage is a helper used to test the output from the examples in this folder.
func TestUsage(usage, command string) {
	args, _ := docopt.ParseArgs(usage, strings.Split(command, " ")[1:], "")

	// Sort the keys of the arguments map
	var keys []string
	for k := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print the argument keys and values
	for _, k := range keys {
		fmt.Printf("%9s %v\n", k, args[k])
	}
	fmt.Println()
}
