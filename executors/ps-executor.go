package executors

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/slc-na/roomnetman-cli/models"
)

var PsExecBinary []byte

func psExec(command models.Command) {
	unpackBinary()
	var waitGroup sync.WaitGroup
	fmt.Printf("[info] Execution will be timed out after 5 seconds\n")
	for _, comp := range command.Computers {
		waitGroup.Add(1)
		cmd := exec.Command("./psexec.exe", fmt.Sprintf("\\\\10.22.%s.%s", comp.Room, comp.Number), "-u", command.User, "-p", command.Password, "-n", "5", "-i", "cmd.exe", "/c", command.Cmd)
		go func(comp models.Computer) {
			runCommand(cmd, fmt.Sprintf("10.22.%s.%s", comp.Room, comp.Number))
			waitGroup.Done()
		}(comp)
	}
	waitGroup.Wait()
	removeBinary()
}

func runCommand(cmd *exec.Cmd, compNum string) {
	err := cmd.Start()
	if err != nil {
		fmt.Printf("[failed] executing script %s\n", compNum)
		return
	}
	errRun := cmd.Wait()
	if errRun != nil {
		fmt.Printf("[failed] finish with err %s: %s\n", compNum, errRun)
		return
	}
	fmt.Printf("[success] executing script %s\n", compNum)
}

func unpackBinary() {
	_ = os.WriteFile("psexec.exe", PsExecBinary, 0755)
}

func removeBinary() {
	os.Remove("psexec.exe")
}
