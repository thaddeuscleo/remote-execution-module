package main

import (
	"embed"

	"github.com/slc-na/ruman-execution-module/database"
	"github.com/slc-na/ruman-execution-module/executors"
)

//go:embed "assets/psexec.exe"
var psexecBinary []byte

//go:embed "assets/computers.csv"
var computerDatabase embed.FS

func NewGoExecution() executors.GoExecution {
	injectEmbedFiles()
	return executors.GoExecution{}
}

func injectEmbedFiles() {
	database.ComputerDatabase = computerDatabase
	executors.PsExecBinary = psexecBinary
}
