package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	cmdSwitch "github.com/erwinhermanto31/quiz_master/cmd"
	"github.com/erwinhermanto31/quiz_master/repo/mysql"
	"github.com/erwinhermanto31/quiz_master/util"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.InitCon()
	mysql.InitMigration()
	fmt.Println("Welcome to Quiz Master! ")
	fmt.Println(" ")

	var command string

	// var num int
	// var i int
	for {
		consoleReader := bufio.NewReader(os.Stdin)
		command, _ = consoleReader.ReadString('\n')
		command = strings.Replace(command, "\n", "", 1)
		fmt.Println(" ")
		cmds := strings.Split(command, " ")

		cmd := cmds[0]
		switch cmd {
		case util.Command[cmd]:
			s := cmdSwitch.NewSwitch(util.Command[cmd])
			err := s.Switch(command)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case "help":
			util.PrintHelp()
			fmt.Println(" ")
		case "exit":
			fmt.Println("===================================")
			os.Exit(0)
		default:
			fmt.Printf("unknown command: %s\n", cmd)
		}
	}

}
