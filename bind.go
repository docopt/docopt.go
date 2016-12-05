package docopt

import (
	"reflect"
	"strings"
	"unicode"
)

/*
Bind binds `args` values to `opt`.
`opt` must be pointer to struct, it returns error if `opt` is not pointer to struct.
`args` will mapped to each exported struct field of `opt`. Examples:
  // Field is ignored by Bind.
  Field int `docopt:"-"`
  // Field mapped from key "--help".
  Field int `docopt:"--help"`
  // Field mapped from key "-h".
  Field int `docopt:"-h"`
  // Field mapped from key "--field".
  Field int
  // F mapped from key "-f"
  F int
*/
func Bind(opt interface{}, args map[string]interface{}) error {
	value := reflect.ValueOf(opt)
	if value.Kind() != reflect.Ptr {
		return newError("'opt' argument is not pointer to struct type")
	}
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return newError("'opt' argument is not pointer to struct type")
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
	for k, v := range args {
		i, ok := indexMap[k]
		if !ok {
			return newError("mapping of \"%s\" is not found in given struct, or an unexported field", k)
		}
		field := value.Field(i)
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			continue
		}
		val := reflect.ValueOf(v)
		if !val.Type().AssignableTo(field.Type()) {
			return newError("value of '%s' is not assignable to '%s' field", k, value.Type().Field(i).Name)
		}
		if !field.CanSet() {
			return newError("'%s' field cannot be set", value.Type().Field(i).Name)
		}
		field.Set(val)
	}
	return nil
}

/*
isUnexportedField returns whether the field is unexported.
isUnexportedField is to avoid the bug in versions older than Go1.3.
See following links:
  https://code.google.com/p/go/issues/detail?id=7247
  http://golang.org/ref/spec#Exported_identifiers
*/
func isUnexportedField(field reflect.StructField) bool {
	return !(field.PkgPath == "" && unicode.IsUpper(rune(field.Name[0])))
}
