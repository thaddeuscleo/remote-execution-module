package models

// Describe the command type and target
type Command struct {
	Cmd       string     // The command that will be executed (optional)
	User      string     // Targeted user that will execute the command
	Password  string     // The User password if any
	Computers []Computer // List of computer that will be executed
	Type      string     // type of operation (run, wake, etc.)
}

// Describe the computer information
// ex. Room 621 computer 12 the ip will be: 10.22.126.112
type Computer struct {
	Room   string // stores the 126 value according to the example
	Number string // stores the 112 value according to the example
}
