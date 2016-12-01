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
