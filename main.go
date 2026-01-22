package main

import (
	"fmt"

	"new_project/config"
	"new_project/database"
	"new_project/service"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	lenPass, symbols := service.GetInput()

	pass := service.GeneratePassword(lenPass, symbols)
	appname := service.GetAppName()
	config.Load()
	db, err := database.Connect(config.Get("DATABASE_URL"), config.Get("DATABASE_PORT"), config.Get("DATABASE_USER"), config.Get("DATABASE_PASS"), config.Get("DATABASE_NAME"))
	if err != nil {
		fmt.Println("error while connect to db!", err)
		return
	}
	fmt.Println("connected to db!")

	database.Save(db, pass, appname)
}
