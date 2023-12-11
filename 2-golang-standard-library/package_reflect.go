package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"50"`
}

type Person struct {
	Name    string `required:"true" max:"50"`
	Email   string `required:"true" max:"50"`
	Address string `required:"false" max:"255"`
}

func readField(value any) {
	// valueType := reflect.TypeOf(value)
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println("Field", structField.Name, "with Type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

// custom validation dengan mengecek struct tag "required" pada struct fields
func IsValid(value any) (result bool) {
	result = true

	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""

			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"John"})
	readField(Person{"John", "john@doe.com", "JL. ABC"})

	person := Person{
		Name:    "Jane",
		Email:   "jane@doe.com",
		Address: "JL. ABC",
	}

	fmt.Println(IsValid(person))
}
