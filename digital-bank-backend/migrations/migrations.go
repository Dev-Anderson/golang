package migrations

import (
	"digital-bank/helpers"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host:127.0.0.1 port=5432 user=postgres dbname=digital-bank password=postgres sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	user := [2]User{
		{Username: "Martin", Email: "margin@martin.com.br"},
		{Username: "Marcos", Email: "marcos@marcos.com.br"},
	}

	for i := 0; i < len(user); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(user[i].Username))
		user := User{Username: user[i].Username, Email: user[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	createAccounts()
}
