package docopt

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
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
// Additionally, Opts.Bind allows you easily populate a struct's fields with the
// values of each option value. See below for examples.
//
// Lastly, you can still treat Opts as a regular map, and do any type checking
// and conversion that you want to yourself. For example:
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

// Bind populates the fields of a given struct with matching option values.
// Each key in Opts will be mapped to an exported field of the struct pointed
// to by `v`, as follows:
//
// 	Field int `docopt:"-"`        // Field is ignored by Bind
// 	Field int `docopt:"--help"`   // Field mapped from key "--help"
// 	Field int `docopt:"-h"`       // Field mapped from key "-h"
// 	Field int                     // Field mapped from key "--field"
// 	F int                         // F mapped from key "-f"
//
func (o Opts) Bind(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return newError("'v' argument is not pointer to struct type")
	}
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return newError("'v' argument is not pointer to struct type")
	}
	typ := value.Type()
	indexMap := make(map[string]int)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if isUnexportedField(field) || field.Anonymous {
			continue
		}
		tag := field.Tag.Get("docopt")
		if tag == "" {
			key := strings.ToLower(field.Name)
			if len(field.Name) == 1 {
				key = "-" + key
			} else {
				key = "--" + key
			}
			indexMap[key] = i
			continue
		}
		for _, t := range strings.Split(tag, ",") {
			indexMap[t] = i
		}
	}
	for k, v := range o {
		i, ok := indexMap[k]
		if !ok {
			if k == "--help" || k == "--version" {
				continue
			}
			return newError("mapping of %q is not found in given struct, or is an unexported field", k)
		}
		field := value.Field(i)
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			continue
		}
		val := reflect.ValueOf(v)
		if !val.Type().AssignableTo(field.Type()) {
			return newError("value of %q is not assignable to %q field", k, value.Type().Field(i).Name)
		}
		if !field.CanSet() {
			return newError("%q field cannot be set", value.Type().Field(i).Name)
		}
		field.Set(val)
	}
	return nil
}

// isUnexportedField returns whether the field is unexported.
// isUnexportedField is to avoid the bug in versions older than Go1.3.
// See following links:
//   https://code.google.com/p/go/issues/detail?id=7247
//   http://golang.org/ref/spec#Exported_identifiers
func isUnexportedField(field reflect.StructField) bool {
	return !(field.PkgPath == "" && unicode.IsUpper(rune(field.Name[0])))
}
