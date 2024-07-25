package demo_command

import (
	"fmt"
	"os/exec"
)

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

type CommandProvider interface {
	HelloCommand() (bool, error)
	GoodNightCommand()
	GoodAfternoonCommand()
}

type CommandService struct {
	CommandProvider
	executor Executor
}

type CmdParams struct {
	FirstName string
	LastName  string
	Gender    string
}

func (cs CommandService) helloCommand(cmdParams CmdParams) (bool, error) {
	commandArgs := []string{"/C", "echo", "hello", cmdParams.FirstName, cmdParams.LastName, cmdParams.Gender}
	output, err := cs.executor.Execute(commandArgs)
	if err != nil {
		fmt.Printf("Error executing hello command: %v\n", err)
		fmt.Printf("Error executing hello command: %s\n", string(output))
		return false, err
	}
	fmt.Println(string(output))
	return true, nil
}
