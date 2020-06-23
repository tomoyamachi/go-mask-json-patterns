package structtag

import (
	"encoding/json"
	"reflect"
)

// Name of the struct tag used in examples
const tagName = "sensitive"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `sensitive:"true" json:"email"`
}

func (u User) MarshalJSON() ([]byte, error) {
	t := reflect.TypeOf(u)
	sensitives := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// tag := field.Tag.Get(tagName)
		// fmt.Printf("%d. %v (%v), val: %s tag: '%v'\n", i+1, field.Name, field.Type.Name(), t.String(), tag)
		if field.Tag.Get(tagName) == "true" {
			sensitives = append(sensitives, field.Name)
		}
	}

	type alias User
	au := alias(u)
	mutable := reflect.ValueOf(&au)
	for _, field := range sensitives {
		f := mutable.Elem().FieldByName(field)
		if f.CanSet() {
			switch f.Kind() {
			case reflect.String:
				f.SetString("***")
			case reflect.Int:
				f.SetInt(99999)
			}
		}
	}
	return json.Marshal(au)
}
