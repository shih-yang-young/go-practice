package examples

type Execer interface {
	Test()
	A()
	B()
	C()
}

type ExecerImpl struct{}

func (e *ExecerImpl) Test() {
	// 原 test 函數的實現
	println("執行 Test 方法...")
}

func (e *ExecerImpl) A() {
	e.Test()
}

func (e *ExecerImpl) B() {
	e.Test()
}

func (e *ExecerImpl) C() {
	e.Test()
}
