package docopt

import (
	"reflect"
	"testing"
)

func TestBind(t *testing.T) {
	const usage = "Usage: prog [-h|--help] [-v] [-f] <command>"
	for i, c := range []struct {
		argv   []string
		expect testoption
	}{
		{[]string{"-v", "test_cmd"}, testoption{
			Command: "test_cmd",
			Help:    false,
			Verbose: true,
			F:       false,
		}},
		{[]string{"-h", "test_cmd"}, testoption{
			Command: "test_cmd",
			Help:    true,
			Verbose: false,
			F:       false,
		}},
		{[]string{"--help", "test_cmd"}, testoption{
			Command: "test_cmd",
			Help:    true,
			Verbose: false,
			F:       false,
		}},
		{[]string{"-f", "test_cmd"}, testoption{
			Command: "test_cmd",
			Help:    false,
			Verbose: false,
			F:       true,
		}},
	} {
		result := testoption{}
		v, err := Parse(usage, c.argv, false, "", false, false)
		if err != nil {
			t.Fatal(err)
		}
		if err := Bind(&result, v); err != nil {
			t.Fatal(err)
		}
		if reflect.DeepEqual(result, c.expect) != true {
			t.Error("testcase:", i, "result:", result, "expect:", c.expect)
		}
	}
}

type testoption struct {
	Command string `docopt:"<command>"`
	Help    bool   `docopt:"-h,--help"`
	Verbose bool   `docopt:"-v"`
	F       bool
}
