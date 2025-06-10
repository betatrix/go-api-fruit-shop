package main

import (
	"log"

	"github.com/betatrix/go-api-fruit-shop/internal/config"
	"github.com/betatrix/go-api-fruit-shop/internal/fruit"
	"github.com/betatrix/go-api-fruit-shop/internal/routes"
	"github.com/betatrix/go-api-fruit-shop/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB    = config.ConnectDB()
	router *gin.Engine = routes.Router(db)
)

func main() {
	err := db.AutoMigrate(&fruit.Fruit{}, &user.User{})
	if err != nil {
		log.Fatal("Erro ao fazer AutoMigrate:", err)
	}

	log.Println("Banco de dados migrado com sucesso!")

	defer config.DisconnectDB(db)
	router.Run(":8080")
}
