package main

import (
	"fmt"
	"log"

	"github.com/erwinhermanto31/quiz_master/repo/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.InitCon()
	mysql.InitMigration()
	var command string
	fmt.Scanf("%s", &command)

	fmt.Printf("%s \n", command)
}
