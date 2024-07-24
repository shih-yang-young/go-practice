// package command

// import (
// 	"fmt"
// 	"os/exec"
// )

// // Execer 是一個執行命令的接口
// type Execer interface {
// 	Command(name string, arg ...string) *exec.Cmd
// }

// // RealExecer 是 Execer 的實現，實際執行命令
// type RealExecer struct{}

// func (r RealExecer) Command(name string, arg ...string) *exec.Cmd {
// 	return exec.Command(name, arg...)
// }

// var execer Execer = RealExecer{}

// func execCommand() bool {
// 	cmd := execer.Command("cmd", "/C", "echo hello I am Louis")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return false
// 	}
// 	fmt.Println(string(output))
// 	return true
// }
