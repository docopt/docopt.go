package docopt

import (
	"reflect"
	"strings"
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
		t.Logf("argv: %v opts: %v", c.argv, v)
		if err != nil {
			t.Fatalf("testcase: %d parse err: %q", i, err)
		}
		if err := v.Bind(&result); err != nil {
			t.Fatalf("testcase: %d bind err: %q", i, err)
		}
		if reflect.DeepEqual(result, c.expect) != true {
			t.Errorf("testcase: %d result: %#v expect: %#v\n", i, result, c.expect)
		}
	}
}

type testTypedOptions struct {
	secret int `docopt:"-s"`

	V       bool
	Number  int16
	Idle    float32
	Pointer uintptr     `docopt:"<ptr>"`
	Ints    []int       `docopt:"<values>"`
	Strings []string    `docopt:"STRINGS"`
	Iface   interface{} `docopt:"IFACE"`
}

func TestBindErrors(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	for i, tc := range []struct {
		usage       string
		command     string
		expectedErr string
	}{
		{
			`Usage: prog [-s]`,
			`prog`,
			`mapping of "-s" is not found in given struct, or is an unexported field`,
		},
		{
			`Usage: prog [--v]`,
			`prog`,
			`mapping of "--v" is not found in given struct, or is an unexported field`,
		},
		{
			`Usage: prog [--number]`,
			`prog`,
			`value of "--number" is not assignable to "Number" field`,
		},
		{
			`Usage: prog [--number=X]`,
			`prog --number=abc`,
			`value of "--number" is not assignable to "Number" field`,
		},
		{
			`Usage: prog <ptr>`,
			`prog 123`,
			`value of "<ptr>" is not assignable to "Pointer" field`,
		},
		{
			`Usage: prog [<values>...]`,
			`prog 123 456`,
			`value of "<values>" is not assignable to "Ints" field`,
		},
		{
			`Usage: prog [-] [IFACE ...]`,
			`prog - 123 456 asd`,
			`mapping of "-" is not found in given struct, or is an unexported field`,
		},
	} {
		argv := strings.Split(tc.command, " ")[1:]
		opts, err := testParser.ParseArgs(tc.usage, argv, "")
		if err != nil {
			t.Fatalf("testcase: %d parse err: %q", i, err)
		}
		var o testTypedOptions
		t.Logf("%#v\n", opts)
		if err := opts.Bind(&o); err != nil {
			if err.Error() != tc.expectedErr {
				t.Fatalf("testcase: %d result: %q expect: %q", i, err.Error(), tc.expectedErr)
			}
		} else {
			t.Fatal("error expected")
		}
	}
}

func TestBindSuccess(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	for i, tc := range []struct {
		usage   string
		command string
	}{
		{
			`Usage: prog [-v]`,
			`prog -v`,
		},
		{
			`Usage: prog [--number=X]`,
			`prog --number=123`,
		},
		{
			`Usage: prog <number>`,
			`prog 123`,
		},
		{
			`Usage: prog NUMBER`,
			`prog 123`,
		},
		{
			`Usage: prog [--idle=X]`,
			`prog --idle=4.1`,
		},
		{
			`Usage: prog [STRINGS ...]`,
			`prog 123 456 asd`,
		},
		{
			`Usage: prog [--help]`,
			`prog --help`,
		},
	} {
		argv := strings.Split(tc.command, " ")[1:]
		opts, err := testParser.ParseArgs(tc.usage, argv, "")
		if err != nil {
			t.Fatalf("testcase: %d parse err: %q", i, err)
		}
		var o testTypedOptions
		t.Logf("%#v\n", opts)
		if err := opts.Bind(&o); err != nil {
			t.Fatalf("testcase: %d error: %q", i, err.Error())
		}
	}
}

func TestBindSimpleStruct(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	opts, err := testParser.ParseArgs("Usage: prog [--number=X]", []string{"--number=123"}, "")
	if err != nil {
		t.Fatal(err)
	}
	var opt struct{ Number int }
	if err := opts.Bind(&opt); err != nil {
		t.Fatal(err)
	}
	if opt.Number != 123 {
		t.Fail()
	}
}

func TestBindToStructWhichAlreadyHasValue(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	opts, err := testParser.ParseArgs("Usage: prog [--number=X]", []string{"--number=123"}, "")
	if err != nil {
		t.Fatal(err)
	}
	var opt = struct{ Number int }{1}
	if err := opts.Bind(&opt); err == nil {
		t.Fatal("error expected")
	}
}

func TestBindDashTag(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	opts, err := testParser.ParseArgs("Usage: prog [-]", []string{"-"}, "")
	if err != nil {
		t.Fatal(err)
	}
	var opt struct {
		Dash bool `docopt:"-"`
	}
	if err := opts.Bind(&opt); err != nil {
		t.Fatal(err)
	}
	if !opt.Dash {
		t.Fail()
	}
}

func TestBindDoubleDashTag(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	opts, err := testParser.ParseArgs("Usage: prog [--]", []string{"--"}, "")
	if err != nil {
		t.Fatal(err)
	}
	var opt struct {
		DoubleDash bool `docopt:"--"`
	}
	if err := opts.Bind(&opt); err != nil {
		t.Fatal(err)
	}
	if !opt.DoubleDash {
		t.Fail()
	}
}

func TestBindHyphenatedTags(t *testing.T) {
	var testParser = &Parser{HelpHandler: NoHelpHandler, SkipHelpFlags: true}
	opts, err := testParser.ParseArgs("Usage: prog --opt-one=N --opt-two=N", []string{"--opt-one", "123", "--opt-two", "234"}, "")
	if err != nil {
		t.Fatal(err)
	}
	var opt struct {
		OptOne string
		OptTwo string
	}
	if err := opts.Bind(&opt); err != nil {
		t.Fatal(err)
	}
	if opt.OptOne != "123" || opt.OptTwo != "234" {
		t.Fail()
	}
}
