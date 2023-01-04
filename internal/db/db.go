package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/k4zb3k/todo-list/pkg/logging"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

func GetDbConnection() (*gorm.DB, error) {
	var (
		host =     getEnvValue("DB_HOST")
		port =     getEnvValue("DB_PORT")
		user =     getEnvValue("DB_USER")
		dbname =   getEnvValue("DB_NAME")
		password = getEnvValue("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
						host, port, user, dbname, password)

	var db *gorm.DB
	db, err := gorm.Open(postgresDriver.Open(conn))
	if err != nil {
		logger.Error(err)
	}
	logger.Infoln("Successful connection to DB")

	return db, nil
}

func getEnvValue(key string) string {
	err := godotenv.Load("/home/k4zb3k/Desktop/score-system/.env")
	if err != nil {
		logger.Error(err)
	}

	return os.Getenv(key)
}