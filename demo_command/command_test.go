package demo_command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommandExecutor 用於模擬 CommandExecutor
type MockCommandExecutor struct {
	mock.Mock
}

// 模擬 Execute 方法
func (m *MockCommandExecutor) Execute(args []string) ([]byte, error) {
	argsCalled := m.Called(args)
	return argsCalled.Get(0).([]byte), argsCalled.Error(1)
}

func TestHelloCommand(t *testing.T) {
	mockExecutor := new(MockCommandExecutor)
	mockExecutor.On("Execute", mock.AnythingOfType("[]string")).Return([]byte("Hello World\n"), nil)

	commandService := CommandService{executor: mockExecutor}

	result, err := commandService.helloCommand(CmdParams{
		FirstName: "Louis",
		LastName:  "Chen",
		Gender:    "sir",
	})
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
	mockExecutor.AssertExpectations(t)
}

func TestHelloCommandError(t *testing.T) {
	mockExecutor := new(MockCommandExecutor)
	mockExecutor.On("Execute", []string{"/C", "echo", "hello", "world"}).Return([]byte("Hello World\n"), errors.New("mock error"))

	commandService := CommandService{executor: mockExecutor}

	result, err := commandService.helloCommand(CmdParams{
		FirstName: "Louis",
		LastName:  "Chen",
		Gender:    "sir",
	})
	assert.Equal(t, false, result)
	assert.Equal(t, errors.New("mock error"), err)
	mockExecutor.AssertExpectations(t)
}
