package database

import (
	"database/sql"
	"fmt"
)

func Save(db *sql.DB, pass string, appName string) bool {
	_, err := db.Exec("INSERT INTO pass (password, app_name) values($1,$2)", pass, appName)
	if err != nil {
		fmt.Println("error while saving!", err)
		return false
	}
	fmt.Println("saved to db!")
	return true
}
