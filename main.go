package main

import (
	"fmt"

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

	db, err := database.Connect("localhost", "5432", "postgres", "nicat123", "go_app")
	if err != nil {
		fmt.Println("error while connect to db!", err)
		return
	}
	fmt.Println("connected to db!")

	database.Save(db, pass, appname)
}
