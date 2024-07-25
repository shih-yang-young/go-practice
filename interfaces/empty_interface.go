package interfaces

import "fmt"

func PrintAnything(v interface{}) {
	fmt.Println(v)
}

func GoEmptyInterface() {
	PrintAnything(42)
	PrintAnything("hello")
	PrintAnything(Person{Name: "Alice"})
}
