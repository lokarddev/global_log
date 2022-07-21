package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Debug bool
	Port  string
)

func InitEnvVariables() error {
	var err error
	if err = envFileLookup(); err != nil {
		Debug, err = checkOSDebug()
		if err != nil {
			if err = loadDefaultEnv(); err != nil {
				log.Println("error loading .env.dev default variables")
			}
		}
	}
	Port = os.Getenv("PORT")
	return err
}

func envFileLookup() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	_, err := godotenv.Read(".env")
	if err == nil {
		log.Println("loaded .env file environment variables")
	}
	return err
}

func checkOSDebug() (bool, error) {
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	return isDebug, err
}

func loadDefaultEnv() error {
	err := godotenv.Load(".env.dev")
	_, err = godotenv.Read(".env.dev")
	if err == nil {
		log.Println(fmt.Sprintf("loaded env variables from .env.dev file"))
	}
	return err
}
