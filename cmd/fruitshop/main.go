package main

import (
	"log"

	"github.com/betatrix/go-api-fruit-shop/internal/config"
	"github.com/betatrix/go-api-fruit-shop/internal/fruit"
	"github.com/betatrix/go-api-fruit-shop/internal/router"
	"github.com/betatrix/go-api-fruit-shop/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB    = config.ConnectDB()
	r  *gin.Engine = router.Router(db)
)

func main() {
	config.GetEnv()
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
