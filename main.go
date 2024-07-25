package main

import (
	"go-practice/examples"
	"go-practice/http"
	"go-practice/interfaces"
)

func main() {
	examples.GoPerson()
	examples.GoIntTree()
	examples.GoEmbedded()
	examples.Go_type_try()

	interfaces.GoImplicitImplement_1()
	interfaces.GoEmptyInterface()
	//demo_command.GoCommand()

	http.GoHttp()
}
