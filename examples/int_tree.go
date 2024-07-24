package examples

import "fmt"

type IntTree struct {
	Val         int
	Left, Right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{Val: val}
	}
	if val < it.Val {
		it.Left = it.Left.Insert(val)
	} else if val > it.Val {
		it.Right = it.Right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.Val:
		return it.Left.Contains(val)
	case val > it.Val:
		return it.Right.Contains(val)
	default:
		return true
	}
}

func GoIntTree() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}
