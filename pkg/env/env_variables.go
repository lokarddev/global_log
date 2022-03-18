package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Debug       bool
	Port        string
	AutoMigrate bool

	DbName       string
	DbUser       string
	DbHost       string
	DbPort       string
	DbPass       string
	DbSsl        string
	DbSchema     string
	MaxCons      int
	WorkersCount int

	RedisDb       string
	RedisHost     string
	RedisPort     string
	RedisUser     string
	RedisProtocol string
	RedisPass     string

	CacheTTL int

	ResultChan string
	LoggerChan string
)

func InitEnvVariables() error {
	var err error
	if err = envFileLookup(); err != nil {
		Debug, err = checkOSDebug()
		if err != nil {
			if err = loadDefaultEnv(); err != nil {
				log.Println("error loading .env.dev default variables")
				return err
			}
		}
	}
	Port = os.Getenv("PORT")
	AutoMigrate, _ = strconv.ParseBool(os.Getenv("AUTOMIGRATE"))

	DbName = os.Getenv("DB_NAME")
	DbUser = os.Getenv("DB_USER")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbPass = os.Getenv("DB_PASS")
	DbSsl = os.Getenv("DB_SSL")
	DbSchema = os.Getenv("DB_SCHEMA")
	MaxCons, _ = strconv.Atoi(os.Getenv("MAX_CONS"))
	WorkersCount, _ = strconv.Atoi(os.Getenv("WORKERS_COUNT"))

	RedisDb = os.Getenv("REDIS_DB")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisUser = os.Getenv("REDIS_USER")
	RedisProtocol = os.Getenv("REDIS_PROTOCOL")
	RedisPass = os.Getenv("REDIS_PASS")

	CacheTTL, _ = strconv.Atoi(os.Getenv("CACHE_TTL"))

	ResultChan = os.Getenv("RESULT_CHAN")
	LoggerChan = os.Getenv("LOGGER_CHAN")

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
