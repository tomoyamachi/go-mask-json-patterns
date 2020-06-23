package sensitive

import (
	"fmt"
	"reflect"
)

// Name of the struct tag used in examples
const tagName = "sensitive"

type User struct {
	Id    int    `sensitive:true`
	Name  string `sensitive:false`
	Email string `sensitive:false`
}

func Mask(obj interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.

	t := reflect.TypeOf(obj)
	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		// Get the field tag value
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v (%v), val: %s tag: '%v'\n", i+1, field.Name, field.Type.Name(), t.String(), tag)
	}

	fmt.Println(obj)
	mutable := reflect.ValueOf(&obj)
	fmt.Println(mutable.Kind())
	if mutable.Kind() == reflect.Ptr {

		fmt.Println(mutable.Elem().Elem())
		mutable.FieldByName("Normal")
		// if f.Kind() == reflect.String {
		// 	fmt.Println(f.CanSet())
		// 	if f.CanSet() {
		// 		f.SetString("update")
		// 	}
		// }
	}
	fmt.Println(obj)
	// f := mutable.FieldByName("Normal").Interface()
	// fmt.Println(f)
	// fmt.Println(obj)
}
