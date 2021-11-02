package structtag_interface

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func Log(v interface{}) ([]byte, error) {
	r := MakeMaskedStruct(v)
	return json.Marshal(r)
}

func MakeMaskedStruct(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}
	rt := reflect.TypeOf(v)
	var rv reflect.Value
	result := map[string]interface{}{}
	switch rt.Kind() {
	case reflect.Struct:
		rv = reflect.ValueOf(&v).Elem().Elem()
	case reflect.Ptr:
		rv = reflect.ValueOf(v).Elem()
		rt = rt.Elem()
	case reflect.Slice:
		log.Println("slice")
	}

	if rt.Kind() != reflect.Struct {
		result["msg"] = v
		return result
	} else if v := checkSpecialStruct(rv); v != nil {
		// special struct set original value
		result["v"] = v
		return result
	}

	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		if isPrivateField(ft.Name) {
			continue
		}
		maskStr := ft.Tag.Get("log")
		jsonTag := ft.Name
		tagOptions := tagOptions("")

		if jt := ft.Tag.Get("json"); jt != "" {
			jsonTag, tagOptions = parseTag(jt)
		}

		if ft.Type.Kind() == reflect.Ptr && rv.Field(i).IsNil() {
			continue
		}
		fv := rv.Field(i)

		// only check log struct tag contains any characters
		if maskStr != "" {
			// set a sensitive character even if the field is struct
			if tagOptions.Contains("omitempty") && isEmptyValue(fv) {
				continue
			}
			result[jsonTag] = maskStr
			continue
		}

		switch ft.Type.Kind() {
		case reflect.Ptr,reflect.Struct:
			result[jsonTag] = MakeMaskedStruct(fv.Interface())
		case reflect.Slice:
			val := []interface{}{}
			for i := 0; i < fv.Len(); i++ {
				inner := fv.Index(i)
				switch inner.Kind() {
				case reflect.Ptr, reflect.Struct:
					val = append(val, MakeMaskedStruct(inner.Interface()))
				default:
					val = append(val, inner.Interface())
				}
			}
			result[jsonTag] = val
		default:
			if tagOptions.Contains("omitempty") && isEmptyValue(fv) {
				continue
			}
			result[jsonTag] = fv.Interface()
		}
	}
	return result
}

func isPrivateField(s string) bool {
	// check first char
	for _, c := range s {
		return !unicode.IsUpper(c)
	}
	return true
}

// return native struct directory
func checkSpecialStruct(rv reflect.Value) interface{} {
	fieldValue := rv.Interface()
	switch fieldValue.(type) {
	case time.Time:
		return fieldValue
	}
	return nil
}

// Copy from https://github.com/golang/go/blob/master/src/encoding/json/tags.go

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
