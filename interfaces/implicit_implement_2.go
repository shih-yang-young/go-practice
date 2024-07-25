package interfaces

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	//
	return "louis"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.Process("data")
}

func GoImplicitImplement_2() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
