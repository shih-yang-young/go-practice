package demo_command

import (
	"fmt"
	"os/exec"
)

type CommandProvider interface {
	HelloCommand() (bool, error)
	GoodNightCommand()
	GoodAfternoonCommand()
}

// 使用Executor 讓Service可以決定要使用哪個Executor ，for測試 or Production
// 一種隱性實作(implemented implicitly)
type Executor interface {
	Execute(args []string) ([]byte, error)
}

type CommandExecutor struct {
}

func (ce CommandExecutor) Execute(args []string) ([]byte, error) {
	cmd := exec.Command("cmd", args...)
	return cmd.CombinedOutput()
}

func (cs CommandService) helloCommand() (bool, error) {
	commandArgs := []string{"/C", "echo", "hello", "world"}
	output, err := cs.executor.Execute(commandArgs)
	if err != nil {
		fmt.Printf("Error executing hello command: %v\n", err)
		return false, err
	}
	fmt.Println(string(output))
	return true, nil
}

type CommandService struct {
	CommandProvider
	executor Executor
}
