package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type env struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func GetDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config = env{}

	config.Host = os.Getenv("host")
	config.Port = os.Getenv("port")
	config.User = os.Getenv("user")
	config.Pass = os.Getenv("pass")
	config.Database = os.Getenv("db")

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Database)
	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	return db
}
