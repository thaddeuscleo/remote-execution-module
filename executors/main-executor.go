package executors

import (
	"github.com/slc-na/roomnetman-cli/models"
)

func ExecuteCommand(command models.Command) {
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
		command.Cmd = "shutdown /s /f /t 0"
		psExec(command)
	case "restart":
		command.Cmd = "shutdown /r /f /t 0"
		psExec(command)
	}
}
