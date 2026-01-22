package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can not load env vars")
	}
}

func Get(key string) string {
	db_name := os.Getenv(key)
	return db_name
}
