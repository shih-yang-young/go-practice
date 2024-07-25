package examples

import "fmt"

type Person1 struct {
	Age int
}

func (p *Person1) AddAge() {
	p.Age++
}

type Person2 struct {
	Age int
}

func (p Person2) AddAge() {
	p.Age++
}

func GoPerson() {
	person1 := Person1{Age: 10}
	person1.AddAge()
	fmt.Println(person1.Age)

	person2 := Person2{Age: 10}
	person2.AddAge()
	fmt.Println(person2.Age)

	person3 := &Person1{Age: 10}
	person3.AddAge()
	fmt.Println(person3.Age)

	person4 := &Person2{Age: 10}
	person4.AddAge()
	fmt.Println(person4.Age)
}
