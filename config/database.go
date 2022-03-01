package config

import (
	"fmt"
	"log"

	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func BuildDBConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     utils.Getenv("DB_HOST", "localhost"),
		Port:     utils.Getenv("DB_PORT", "3306"),
		User:     utils.Getenv("DB_USER", "root"),
		Password: utils.Getenv("DB_PASSWORD", ""),
		DbName:   utils.Getenv("DB_NAME", "blog"),
	}
}

func DatabaseUri(dbConfig *DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}

func ConnectMysql() *gorm.DB {
	var dsn string = DatabaseUri(BuildDBConfig())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Failed")
		log.Println(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	return db
}