package interfaces

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func GoImplicitImplement_1() {
	var s Speaker
	p := Person{Name: "Alice"}
	s = p

	fmt.Println(s.Speak())
}
