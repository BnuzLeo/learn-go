package operator

import (
	"fmt"
	"reflect"
)

type Person struct {
	Address Address
	Name    string
	Age     int
}

type Address struct {
	Name string
}

func UpdatePersonInfo(person *Person, name string) {
	fmt.Println(reflect.TypeOf(person))
	fmt.Println(reflect.TypeOf(*person))
	person.Address.Name = name
}
