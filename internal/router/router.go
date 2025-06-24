package router

import (
	"github.com/betatrix/go-api-fruit-shop/internal/domain/fruit"
	"github.com/betatrix/go-api-fruit-shop/internal/domain/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	fruit.RegisterFruitRoutes(route, db)
	user.RegisterUserRoutes(route, db)

	return route
}
