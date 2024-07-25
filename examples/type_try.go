package examples

import "fmt"

type TwoInt func(a int, b int) int

func (ti TwoInt) DefaultAdd(c int, d int) {
	ti(c, d)
}

func add(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func Go_type_try() {
	var twoInt TwoInt
	twoInt = add
	fmt.Println(twoInt(3, 4)) // Output: 7
	twoInt = minus
	fmt.Println(twoInt(3, 4)) // Output: -1
	twoInt = times
	fmt.Println(twoInt(3, 4)) // Output: 12

	twoInt.DefaultAdd(5, 6)
}
