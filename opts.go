package docopt

import (
	"fmt"
	"strconv"
)

func errKey(key string) error {
	return fmt.Errorf("no such key: %s", key)
}
func errType(key string) error {
	return fmt.Errorf("key: %s failed type conversion", key)
}
func errStrconv(key string, convErr error) error {
	return fmt.Errorf("key: %s failed type conversion: %s", key, convErr)
}

// Opts is a map of command line options to their values, with some convenience
// methods for value type conversion (bool, float64, int, string). For example,
// to get an option value as an int:
//
// 	opts, _ := docopt.ParseDoc("Usage: sleep <seconds>")
// 	secs, _ := opts.Int("<seconds>")
//
// You can still treat the Opts as a regular map, and do any type checking and
// conversion that you want to yourself. For example:
//
// 	if s, ok := opts["<binary>"].(string); ok {
// 		if val, err := strconv.ParseUint(s, 2, 64); err != nil { ... }
// 	}
//
// Note that any non-boolean option / flag will have a string value in the
// underlying map.
type Opts map[string]interface{}

func (o Opts) String(key string) (s string, err error) {
	v, ok := o[key]
	if !ok {
		err = errKey(key)
		return
	}
	s, ok = v.(string)
	if !ok {
		err = errType(key)
	}
	return
}

func (o Opts) Bool(key string) (b bool, err error) {
	v, ok := o[key]
	if !ok {
		err = errKey(key)
		return
	}
	b, ok = v.(bool)
	if !ok {
		err = errType(key)
	}
	return
}

func (o Opts) Int(key string) (i int, err error) {
	s, err := o.String(key)
	if err != nil {
		return
	}
	i, err = strconv.Atoi(s)
	if err != nil {
		err = errStrconv(key, err)
	}
	return
}

func (o Opts) Float64(key string) (f float64, err error) {
	s, err := o.String(key)
	if err != nil {
		return
	}
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		err = errStrconv(key, err)
	}
	return
}
