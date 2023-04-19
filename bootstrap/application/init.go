package application

import (
	"QuizApp/config"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func SetOutputLog() {

	var logFileName string

	logChannel := os.Getenv("LOG_CHANNEL")
	switch logChannel {
	case string(config.DAILY):
		timeNow := time.Now()
		dateTimeString := fmt.Sprintf("%d-%d-%d", timeNow.Year(), timeNow.Month(), timeNow.Day())
		logFileName = fmt.Sprintf("log-%s.log", dateTimeString)
	default:
		logFileName = "log.log"
	}

	filePath := fmt.Sprintf("./storage/log/%s", logFileName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Printf("We are logging in Golang Based on LOG CHANNEL: %s!", logChannel)
}

func LoadingEnvironmentVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Loading Environment file errors ocuured! Err: %s", err)
	}
}
