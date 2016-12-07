package docopt

import (
	"reflect"
	"testing"
)

func TestOptsUsage(t *testing.T) {
	usage := "Usage: sleep <seconds> [--now]"
	var opts Opts

	opts, _ = ParseArgs(usage, []string{"10"}, "")
	i, err := opts.Int("<seconds>")
	if err != nil || !reflect.DeepEqual(i, int(10)) {
		t.Fail()
	}
	f, err := opts.Float64("<seconds>")
	if err != nil || !reflect.DeepEqual(f, float64(10)) {
		t.Fail()
	}

	opts, _ = ParseArgs(usage, []string{"ten"}, "")
	s, err := opts.String("<seconds>")
	if err != nil || !reflect.DeepEqual(s, string("ten")) {
		t.Fail()
	}

	opts, _ = ParseArgs(usage, []string{"10", "--now"}, "")
	b, err := opts.Bool("--now")
	if err != nil || !reflect.DeepEqual(b, true) {
		t.Fail()
	}
}

func TestOptsErrors(t *testing.T) {
	usage := "Usage: sleep <seconds> [--now]"
	var opts Opts
	var err error

	opts, _ = ParseArgs(usage, []string{"ten!"}, "")

	_, err = opts.Int("<seconds>") // errStrconv
	if err == nil {
		t.Fail()
	}
	_, err = opts.Float64("<seconds>") // errStrconv
	if err == nil {
		t.Fail()
	}

	_, err = opts.Bool("<seconds>") // errType
	if err == nil {
		t.Fail()
	}
	_, err = opts.String("--now") // errType
	if err == nil {
		t.Fail()
	}
	_, err = opts.Int("--now") // errType
	if err == nil {
		t.Fail()
	}
	_, err = opts.Float64("--now") // errType
	if err == nil {
		t.Fail()
	}

	_, err = opts.Int("<missing>") // errKey
	if err == nil {
		t.Fail()
	}
	_, err = opts.Float64("<missing>") // errKey
	if err == nil {
		t.Fail()
	}
	_, err = opts.Bool("<missing>") // errKey
	if err == nil {
		t.Fail()
	}
	_, err = opts.String("<missing>") // errKey
	if err == nil {
		t.Fail()
	}
}

func TestOptsBind(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
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
		v, err := testParser.ParseArgs(usage, c.argv, "")
		if err != nil {
			t.Fatal(err)
		}
		if err := v.Bind(&result); err != nil {
			t.Fatal(err)
		}
		if reflect.DeepEqual(result, c.expect) != true {
			t.Errorf("testcase: %#v result: %#v expect: %#v\n", i, result, c.expect)
		}
	}
}

type testoption struct {
	Command string `docopt:"<command>"`
	Help    bool   `docopt:"-h,--help"`
	Verbose bool   `docopt:"-v"`
	F       bool
}
