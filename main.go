package main

import (
	"embed"
	"fmt"

	"github.com/slc-na/roomnetman-cli/database"
	"github.com/slc-na/roomnetman-cli/executors"
	"github.com/slc-na/roomnetman-cli/utils"
)

//go:embed "assets/psexec.exe"
var psexecBinary []byte

//go:embed "assets/computers.csv"
var computerDatabase embed.FS

func injectEmbedFiles() {
	database.ComputerDatabase = computerDatabase
	executors.PsExecBinary = psexecBinary
}

func main() {
	fmt.Printf("| Minion |\n\n")
	injectEmbedFiles()

	// get cli arguments
	cmd := utils.GetUserArgs()
	executors.ExecuteCommand(cmd)
}
