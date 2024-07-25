package main

import (
	"go-practice/demo_command"
	"go-practice/examples"
	"go-practice/interfaces"
)

func main() {
	execer := &examples.ExecerImpl{}
	execer.A()
	execer.B()
	execer.C()
	// Person Example
	examples.GoPerson()
	examples.GoIntTree()
	examples.GoEmbedded()

	interfaces.GoImplicitImplement_1()
	interfaces.GoEmptyInterface()
	demo_command.GoCommand()
}
