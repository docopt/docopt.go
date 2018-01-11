package main

import (
	"github.com/docopt/docopt-go"
	"reflect"
	"testing"
)

var usage = `Usage:
  nettool tcp <host> <port> [--timeout=<seconds>]
  nettool serial <port> [--baud=9600] [--timeout=<seconds>]
  nettool -h | --help | --version`

// List of test cases
var usageTestTable = []struct {
	argv      []string    // Given command line args
	validArgs bool        // Are they supposed to be valid?
	opts      docopt.Opts // Expected options parsed
}{
	{
		[]string{"tcp", "myhost.com", "8080", "--timeout=20"},
		true,
		docopt.Opts{
			"--baud":    nil,
			"--help":    false,
			"--timeout": "20",
			"--version": false,
			"-h":        false,
			"<host>":    "myhost.com",
			"<port>":    "8080",
			"serial":    false,
			"tcp":       true,
		},
	},
	{
		[]string{"serial", "1234", "--baud=14400"},
		true,
		docopt.Opts{
			"--baud":    "14400",
			"--help":    false,
			"--timeout": nil,
			"--version": false,
			"-h":        false,
			"<host>":    nil,
			"<port>":    "1234",
			"serial":    true,
			"tcp":       false,
		},
	},
	{
		[]string{"foo", "bar", "dog"},
		false,
		docopt.Opts{},
	},
}

func TestUsage(t *testing.T) {
	for _, tt := range usageTestTable {
		validArgs := true
		parser := &docopt.Parser{
			HelpHandler: func(err error, usage string) {
				if err != nil {
					validArgs = false // Triggered usage, args were invalid.
				}
			},
		}
		opts, err := parser.ParseArgs(usage, tt.argv, "")
		if validArgs != tt.validArgs {
			t.Fail()
		}
		if tt.validArgs && err != nil {
			t.Fail()
		}
		if tt.validArgs && !reflect.DeepEqual(opts, tt.opts) {
			t.Errorf("result (1) doesn't match expected (2) \n%v \n%v", opts, tt.opts)
		}
	}
}
