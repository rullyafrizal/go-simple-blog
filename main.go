package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rullyafrizal/go-simple-blog/config"
	"github.com/rullyafrizal/go-simple-blog/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env")
	}

	db := config.ConnectMysql()
	sqlDb, err := db.DB()

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = sqlDb.Ping()

	if err != nil {
		log.Fatal("Error pinging database")
	}

	defer sqlDb.Close()

	routes.SetupRouter(db)
}