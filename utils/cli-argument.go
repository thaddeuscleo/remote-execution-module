package utils

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thaddeuscleo/remote-execution-module/models"
)

func GetUserArgs() models.Command {
	var userCmd models.Command

	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	wakeCmd := flag.NewFlagSet("wake", flag.ExitOnError)
	shutdownCmd := flag.NewFlagSet("shutdown", flag.ExitOnError)
	restartCmd := flag.NewFlagSet("restart", flag.ExitOnError)
	deepCmd := flag.NewFlagSet("deep", flag.ExitOnError)
	undeepCmd := flag.NewFlagSet("undeep", flag.ExitOnError)

	// computers := []string{"10.22.110.101", "10.22.110.102", "10.22.110.103", "10.22.110.104", "10.22.110.105"}
	computers := []string{}

	runUser := runCmd.String("u", "root", "By default the user is set to root")
	runPassword := runCmd.String("p", "", "The computer(s) password")
	runRoom := runCmd.String("r", "", "The computer(s) room, ex. 626")
	runExcutable := runCmd.String("x", "", "The command That will be executed")

	wakeRoom := wakeCmd.String("r", "", "The computer(s) room, ex. 626")

	shutdownRoom := shutdownCmd.String("r", "", "The computer(s) room, ex. 626")
	shutdownPassword := shutdownCmd.String("p", "", "The computer(s) password")

	restartRoom := restartCmd.String("r", "", "The computer(s) room, ex. 626")
	restartPassword := restartCmd.String("p", "", "The computer(s) password")

	deepRoom := deepCmd.String("r", "", "The computer(s) room, ex. 626")
	deepPassword := deepCmd.String("p", "", "The computer(s) password")

	undeepRoom := undeepCmd.String("r", "", "The computer(s) room, ex. 626")
	undeepPassword := undeepCmd.String("p", "", "The computer(s) password")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'run', 'wake', 'shutdown', 'restart', 'deep', 'undeep', 'log' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "run":
		runCmd.Parse(os.Args[2:])

		if *runPassword == "" {
			log.Fatal("[Error] Password Is Required (ex. -p pass)")
		}

		if *runRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		if *runExcutable == "" {
			log.Fatal("[Error] Command Is Required")
		}

		userCmd = models.Command{
			Type:      "run",
			User:      *runUser,
			Password:  *runPassword,
			Cmd:       *runExcutable,
			Computers: computers,
		}
	case "wake":
		wakeCmd.Parse(os.Args[2:])
		if *wakeRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		userCmd = models.Command{
			Type:      "wake",
			Computers: computers,
		}
	case "shutdown":
		shutdownCmd.Parse(os.Args[2:])
		if *shutdownRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		userCmd = models.Command{
			Type:      "shutdown",
			User:      "root",
			Password:  *shutdownPassword,
			Computers: computers,
		}
	case "restart":
		restartCmd.Parse(os.Args[2:])
		if *restartRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		userCmd = models.Command{
			Type:      "restart",
			User:      "root",
			Password:  *restartPassword,
			Computers: computers,
		}
	case "deep":
		deepCmd.Parse(os.Args[2:])
		if *deepRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		userCmd = models.Command{
			Type:      "deep",
			User:      "root",
			Password:  *deepPassword,
			Computers: computers,
		}
	case "undeep":
		undeepCmd.Parse(os.Args[2:])
		if *undeepRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}

		userCmd = models.Command{
			Type:      "undeep",
			User:      "root",
			Password:  *undeepPassword,
			Computers: computers,
		}
	default:
		flag.PrintDefaults()
		fmt.Println("expected 'run' or 'wake' subcommands")
		os.Exit(1)
	}
	return userCmd
}
