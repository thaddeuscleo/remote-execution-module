package executors

import "github.com/thaddeuscleo/remote-execution-module/models"

type goExecution models.Command

func (command goExecution) ExecuteCommand() {
	switch command.Type {
	case "run":
		psExec(command)
	case "wake":
		wakeExec(command)
	case "deep":
		command.Cmd = "uwfmgr filter enable"
		psExec(command)
	case "undeep":
		command.Cmd = "uwfmgr filter disable"
		psExec(command)
	case "shutdown":
		command.Cmd = "shutdown.exe /s /f /t 0"
		psExec(command)
	case "restart":
		command.Cmd = "shutdown /r /f /t 0"
		psExec(command)
	}
}

func RunExecutionCommand(command models.Command) goExecution {
	return goExecution{
		User:      command.User,
		Password:  command.Password,
		Cmd:       command.Cmd,
		Computers: command.Computers,
		Type:      command.Type,
	}
}
