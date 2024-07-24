package example

type IntTree struct {
	Val         int
	Left, Right *IntTree
}

func (it *IntTree) IntTree(val int) *IntTree {
	if it == nil {
		return &IntTree{Val: val}
	}
	if val < it.Val {
		it.Left = it.Left.IntTree(val)
	} else if val > it.Val {
		it.Right = it.Right.IntTree(val)
	}
	return it
}
