package reflect

import (
	"errors"
	"fmt"
)

type Number int

type Content string

type Person struct {
	Address `json:"address"`
	Name    Content `json:"name"`
	Age     Number  `json:"age"`
	sex     string  `json:"sex"`
}

type Address struct {
	Country  string `json:"country"`
	Province string `json:"Province"`
	Street   string `json:"street"`
	door     int    `json:"door"`
}

type A int

type B struct {
	A
}

func (A) av() {

}

func (*A) ap() {

}

func (B) bv() {

}

func (*B) bp() {

}

func (p Person) Run(duration int, length int) error {
	fmt.Println(p.Name, "run", length, "m, use", duration)
	return errors.New("run fail")
}

func (p Person) RunLongDuration(duration int, lengths ...int) error {
	for _, length := range lengths {
		fmt.Println(p.Name, "run", length, "m, use", duration)
	}
	return errors.New("run fail")
}
