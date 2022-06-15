package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	host     string
	port     string
	user     string
	pass     string
	database string
}

func GetDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config = env{}

	config.host = os.Getenv{"host"}
	config.port = os.Getenv{"port"}
	config.user = os.Getenv{"user"}
	config.pass = os.Getenv{"pass"}
	config.database = os.Getenv{"db"}

	info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.host, config.user, config.pass, config.database)
	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	return db
}
