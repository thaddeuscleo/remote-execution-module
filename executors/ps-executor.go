// This package contains executor file
package executors

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/thaddeuscleo/remote-execution-module/utils"
)

//go:embed "bin/psexec.exe"
var psexecBinary []byte

func psExec(command goExecution) {
	unpackBinary()
	var waitGroup sync.WaitGroup
	fmt.Printf("[info] Execution will be timed out after 5 seconds\n")
	for _, comp := range command.Computers {
		waitGroup.Add(1)
		cmd := exec.Command("./psexec.exe", fmt.Sprintf("\\\\%s", comp), "-u", command.User, "-p", command.Password, "-n", "5", "-i", "cmd.exe", "/c", command.Cmd)
		go func(comp string) {
			runCommand(cmd, comp)
			waitGroup.Done()
		}(comp)
	}
	waitGroup.Wait()
	removeBinary()
}

func runCommand(cmd *exec.Cmd, compNum string) {
	err := cmd.Start()
	if err != nil {
		content := fmt.Sprintf("executing script %s\n", compNum)
		utils.LogError(content)
		return
	}
	errRun := cmd.Wait()
	if errRun != nil {
		content := fmt.Sprintf("finish with err %s: %s\n", compNum, errRun)
		utils.LogError(content)
		return
	}
	content := fmt.Sprintf("executing script %s\n", compNum)
	utils.LogSuccess(content)
}

func unpackBinary() {
	_ = os.WriteFile("psexec.exe", psexecBinary, 0755)
}

func removeBinary() {
	os.Remove("psexec.exe")
}
