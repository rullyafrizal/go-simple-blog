package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rullyafrizal/go-simple-blog/config"
	"github.com/rullyafrizal/go-simple-blog/routes"
	"github.com/rullyafrizal/go-simple-blog/seeds"
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

	err = seeds.Seed(db)

	if err != nil {
		log.Fatal("Error seeding database")
	}

	defer sqlDb.Close()

	routes.SetupRouter(db)
}
