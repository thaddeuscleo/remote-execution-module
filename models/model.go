package models

type GoExecution Command

// Describe the command type and target
type Command struct {
	Cmd       string   // The command that will be executed (optional)
	User      string   // Targeted user that will execute the command
	Password  string   // The User password if any
	Computers []string // List of computer that will be executed
	Type      string   // type of operation (run, wake, etc.)
}
