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

type testOptions struct {
	Command string `docopt:"<command>"`
	Help    bool   `docopt:"-h,--help"`
	Verbose bool   `docopt:"-v"`
	F       bool
}

func TestOptsBind(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	const usage = "Usage: prog [-h|--help] [-v] [-f] <command>"
	for i, c := range []struct {
		argv   []string
		expect testOptions
	}{
		{[]string{"-v", "test_cmd"}, testOptions{
			Command: "test_cmd",
			Help:    false,
			Verbose: true,
			F:       false,
		}},
		{[]string{"-h", "test_cmd"}, testOptions{
			Command: "test_cmd",
			Help:    true,
			Verbose: false,
			F:       false,
		}},
		{[]string{"--help", "test_cmd"}, testOptions{
			Command: "test_cmd",
			Help:    true,
			Verbose: false,
			F:       false,
		}},
		{[]string{"-f", "test_cmd"}, testOptions{
			Command: "test_cmd",
			Help:    false,
			Verbose: false,
			F:       true,
		}},
	} {
		result := testOptions{}
		v, err := testParser.ParseArgs(usage, c.argv, "")
		if err != nil {
			t.Fatal(err)
		}
		if err := v.Bind(&result); err != nil {
			t.Fatal(err)
		}
		if reflect.DeepEqual(result, c.expect) != true {
			t.Errorf("testcase: %d result: %#v expect: %#v\n", i, result, c.expect)
		}
	}
}

type testTypedOptions struct {
	seconds int     `docopt:"-s"`
	Size    int     `docopt:"-n"`
	Idle    float64 `docopt:"-i"`
	Command string  `docopt:"<command>"`
	Force   bool    `docopt:"--force"`
	Verbose bool    `docopt:"--verbose"`
}

func TestBindErrors(t *testing.T) {
	const usage = "Usage: prog [-s]"
	for i, tc := range []struct {
		usage       string
		expectedErr string
	}{
		{
			`Usage: prog [-s]`,
			`mapping of "-s" is not found in given struct, or is an unexported field`,
		},
		{
			`Usage: prog [-n] [-i] [<command>]`,
			`value of "-n" is not assignable to "Size" field`,
		},
	} {
		opts, err := ParseArgs(tc.usage, []string{}, "")
		if err != nil {
			t.Fatal(err)
		}
		var o testTypedOptions
		if err := opts.Bind(&o); err != nil {
			if err.Error() != tc.expectedErr {
				t.Errorf("testcase: %d result: %q expect: %q", i, err.Error(), tc.expectedErr)
			}
		} else {
			t.Fatal("error expected")
		}
	}
}
