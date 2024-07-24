package main

import (
	"go-practice/examples"
)

func main() {
	// Person Example
	examples.GoPerson()
	examples.GoIntTree()

	execer := &examples.ExecerImpl{}
	execer.A()
	execer.B()
	execer.C()

}
