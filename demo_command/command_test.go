package demo_command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockCommandExecutor 用於模擬 CommandExecutor
type MockCommandExecutor struct {
	mock.Mock
}

// 模擬 Execute 方法
func (m *MockCommandExecutor) Execute(args []string) ([]byte, error) {
	fmt.Println("I am argsargsargs", args)
	argsCalled := m.Called(args)
	fmt.Println("I am argsCalledargsCalledargsCalled", argsCalled)

	return argsCalled.Get(0).([]byte), argsCalled.Error(1)
}

// 測試 CommandService 的 HelloCommand 方法
func TestHelloCommand(t *testing.T) {
	mockExecutor := new(MockCommandExecutor)
	mockExecutor.On("Execute", []string{"/C", "echo Hello World"}).Return([]byte("Hello World\n"), nil)

	commandService := CommandService{executor: mockExecutor}

	commandService.helloCommand()

	mockExecutor.AssertExpectations(t)
}
