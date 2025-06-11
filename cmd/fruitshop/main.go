package main

import (
	"log"

	"github.com/betatrix/go-api-fruit-shop/internal/config"
	"github.com/betatrix/go-api-fruit-shop/internal/domain/fruit"
	"github.com/betatrix/go-api-fruit-shop/internal/domain/user"
	"github.com/betatrix/go-api-fruit-shop/internal/router"
)

func main() {
	config.GetEnv()

	db := config.ConnectDB()
	r := router.Router(db)

	defer config.DisconnectDB(db)

	err := db.AutoMigrate(&fruit.Fruit{}, &user.User{})
	if err != nil {
		log.Fatal("Error during AutoMigrate:", err)
	}
	log.Println("Database migrated successfully!")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	log.Println("Connected to port: 8080")
}
