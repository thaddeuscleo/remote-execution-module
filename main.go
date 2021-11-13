package main

import (
	"io/ioutil"

	"github.com/thaddeuscleo/remote-execution-module/database"
	"github.com/thaddeuscleo/remote-execution-module/executors"
	"github.com/thaddeuscleo/remote-execution-module/utils"
)

func main() {
	file, _ := ioutil.ReadFile("computers.csv")
	database.SetComputersDatabase(file)
	cmd := utils.GetUserArgs()
	executors.RunExecutionCommand(cmd).ExecuteCommand()
}
