package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

type testStruct struct {
	A string `json:"name_of_a"`
	B int    `json:"holly_gabrielle_dwinell_cole"`
}

func main() {
	t := testStruct{A: "test", B: 10}
	s, err := structTag(t)
	if err != nil {
		fmt.Print(err)
	}
	// s = s.(testStruct)
	printStructTags(s)
}

func structTag(i interface{}) (interface{}, error) {
	v := reflect.ValueOf(i)
	// check if it's a pointer
	if v.Kind() == reflect.Ptr {
		// get the pointer element
		v = v.Elem()
	}
	// test to make sure we have a struct
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("Type is not a struct or a pointer to a struct")
	}
	// get type
	t := v.Type()
	structFields := []reflect.StructField{}
	// loop over all the fields in the struct
	for i := 0; i < t.NumField(); i++ {
		// get field type
		f := t.Field(i)
		// test if exported
		// if f.PkgPath != "" {
		// 	continue
		// }
		// replace the json tag with a proper name
		structFields = append(structFields, replaceTag(f))
	}
	var structType reflect.Type
	structType = reflect.StructOf(structFields)
	return reflect.Zero(structType).Interface(), nil
}

func replaceTag(f reflect.StructField) reflect.StructField {
	sf := reflect.StructField{
		Name:      f.Name,
		PkgPath:   f.PkgPath,
		Type:      f.Type,
		Offset:    f.Offset,
		Index:     f.Index,
		Anonymous: f.Anonymous,
	}
	val, ok := f.Tag.Lookup("json")
	if !ok {
		sf.Tag = reflect.StructTag(`json:"-"`)
		return sf
	}
	opts := strings.Split(val, ",")
	opts[0] = strcase.ToLowerCamel(opts[0])
	opt := strings.Join(opts, ",")
	sf.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, opt))
	return sf
}

func printStructTags(i interface{}) {
	v := reflect.ValueOf(i)
	// check if it's a pointer
	if v.Kind() == reflect.Ptr {
		// get the pointer element
		v = v.Elem()
	}
	// test to make sure we have a struct
	if v.Kind() != reflect.Struct {
		fmt.Println("Type is not a struct or a pointer to a struct")
	}
	// get type
	t := v.Type()
	// loop over all the fields in the struct
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag

		fmt.Printf("Name: %s, Tag: %s\n", name, tag)
	}
}
