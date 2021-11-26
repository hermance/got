package main

import (
	"api/src/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	load_env()
	db := connect_to_db()
	migrate(db)
	listen(db)
}

func listen(db *gorm.DB) {
	r := gin.Default()
	routes.Bind_routes(r, db)
	r.Run()
}

func connect_to_db() *gorm.DB {
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func load_env() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Organization{}, &User{}, &UserConfig{}, &WorkDay{})
}
