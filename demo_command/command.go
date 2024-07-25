package demo_command

import (
	"fmt"
	"os/exec"
)

type CommandProvider interface {
	HelloCommand()
	GoodNightCommand()
	GoodAfternoonCommand()
}

type Executor interface {
	Execute(args []string) ([]byte, error)
}

type CommandExecutor struct {
}

func (ce CommandExecutor) Execute(args []string) ([]byte, error) {
	cmd := exec.Command("cmd", args...)
	return cmd.CombinedOutput()
}

func (cs CommandService) helloCommand() {
	commandArgs := []string{"/C", "echo Hello World"}
	output, _ := cs.executor.Execute(commandArgs)
	fmt.Println(string(output))
}

type CommandService struct {
	CommandProvider
	executor Executor
}
