package executors

import "github.com/slc-na/ruman-execution-module/models"

type GoExecution models.GoExecution

func (command GoExecution) ExecuteCommand() {
	switch command.Type {
	case "run":
		psExec(command)
	case "wake":
		wakeExec(command)
	case "log":
		// TODO: Log Executor
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

func (execCommand GoExecution) SetExecutionCommand(command models.Command) GoExecution {
	return GoExecution{
		User:      command.User,
		Password:  command.Password,
		Cmd:       command.Cmd,
		Computers: command.Computers,
		Type:      command.Type,
	}
}
