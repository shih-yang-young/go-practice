// package command

// import (
// 	"os/exec"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockExecer 是一個實現 Execer 接口的 mock 結構體
// type MockExecer struct {
// 	mock.Mock
// }

// func (m *MockExecer) Command(name string, arg ...string) *exec.Cmd {
// 	args := m.Called(name, arg)
// 	return args.Get(0).(*exec.Cmd)
// }

// // MockCmd 是一個用來 mock exec.Cmd 的結構體
// type MockCmd struct {
// 	exec.Cmd
// 	mock.Mock
// }

// func (m *MockCmd) Output() ([]byte, error) {
// 	args := m.Called()
// 	return args.Get(0).([]byte), args.Error(1)
// }

// // fakeExecCommand 用來替代 exec.Command 返回 MockCmd
// func fakeExecCommand(name string, args ...string) *exec.Cmd {
// 	mockCmd := new(MockCmd)
// 	mockCmd.On("Output").Return("hello I am Louis", nil)
// 	return &mockCmd.Cmd
// }

// func TestExecCommand(t *testing.T) {
// 	mockExecer := new(MockExecer)
// 	mockExecer.On("Command", "cmd", "/C", "echo hello I am Louis").Return(fakeExecCommand("cmd", "/C", "echo hello I am Louis"))

// 	// 使用 mockExecer 替換原來的 execer
// 	execer = mockExecer

// 	result := execCommand()
// 	assert.True(t, result)
// 	mockExecer.AssertExpectations(t)
// }