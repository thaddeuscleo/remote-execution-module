package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/slc-na/roomnetman-cli/models"
)

func GetUserArgs() models.Command {
	var userCmd models.Command

	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	wakeCmd := flag.NewFlagSet("wake", flag.ExitOnError)
	shutdownCmd := flag.NewFlagSet("shutdown", flag.ExitOnError)

	runUser := runCmd.String("u", "root", "By default the user is set to root")
	runPassword := runCmd.String("p", "", "The computer(s) password")
	runComputer := runCmd.String("c", "1-41", "By the default it is set to computer 1 - 41. You can change for ex. 1-3 / 21(one pc)")
	runRoom := runCmd.String("r", "", "The computer(s) room, ex. 626")
	runExcutable := runCmd.String("x", "", "The command That will be executed")

	wakeRoom := wakeCmd.String("r", "", "The computer(s) room, ex. 626")
	wakeComputer := wakeCmd.String("c", "1-41", "By the default it is set to computer 1 - 41. You can change for ex. 1-3 / 21(one pc)")

	shutdownRoom := shutdownCmd.String("r", "", "The computer(s) room, ex. 626")
	shutdownComputer := shutdownCmd.String("c", "1-41", "By the default it is set to computer 1 - 41")
	shutdownPassword := shutdownCmd.String("p", "", "The computer(s) password")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'run', 'wake', 'shutdown' subcommands")
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

		roomCvt := parseRoom(runRoom)
		computers := parseComputer(runComputer, roomCvt)

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

		roomCvt := parseRoom(wakeRoom)
		computers := parseComputer(wakeComputer, roomCvt)

		userCmd = models.Command{
			Type:      "wake",
			Computers: computers,
		}
	case "shutdown":
		shutdownCmd.Parse(os.Args[2:])
		if *shutdownRoom == "" {
			log.Fatal("[Error] Room Is Required")
		}
		roomCvt := parseRoom(shutdownRoom)
		computers := parseComputer(shutdownComputer, roomCvt)
		userCmd = models.Command{
			Type:      "shutdown",
			Password:  *shutdownPassword,
			Computers: computers,
		}
	default:
		flag.PrintDefaults()
		fmt.Println("expected 'run' or 'wake' subcommands")
		os.Exit(1)
	}
	return userCmd
}

func parseRoom(room *string) string {
	roomCvt := *room

	if roomCvt[:1] == "6" {
		roomCvt = strings.Replace(roomCvt, "6", "1", 1)
	} else if roomCvt[:1] == "7" {
		roomCvt = strings.Replace(roomCvt, "7", "2", 1)
	} else {
		log.Fatal("[Error] Invalid room")
	}

	return roomCvt
}

func parseComputer(computer *string, room string) []models.Computer {
	var computers []models.Computer
	if strings.Contains(*computer, "-") {
		split := strings.Split(*computer, "-")

		if len(split) != 2 {
			log.Fatal("[Error] Computer range is invalid")
		}

		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])

		for i := start; i <= end; i++ {
			computers = append(computers, models.Computer{
				Room:   room,
				Number: fmt.Sprintf("1%02d", i),
			})
		}
	} else {
		comp, err := strconv.Atoi(*computer)

		if err != nil {
			log.Fatal("[Error] Invalid computer")
		}

		computers = append(computers, models.Computer{
			Room:   room,
			Number: fmt.Sprintf("1%02d", comp),
		})
	}

	return computers
}
