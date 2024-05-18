package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/post-app.git/backend/models"
	"github.com/iqbalmahad/post-app.git/backend/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/crud_post_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	db.AutoMigrate(&models.Post{})

	return db
}

func main() {
	app := fiber.New()

	db := initDatabase()
	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
