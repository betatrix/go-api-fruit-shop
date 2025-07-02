package fruit

import (
	"github.com/betatrix/go-api-fruit-shop/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFruitRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewFruitRepository(db)
	service := NewFruitService(repo)
	handler := NewFruitHandler(service)

	//user routes
	userGroup := r.Group("/")
	userGroup.Use(
		middlewares.AuthMiddleware(),
		middlewares.AuthorizationRole("admin", "user"),
	)

	userGroup.GET("/fruits/:id", handler.GetFruitbyID)
	userGroup.GET("/fruits", handler.GetAllFruits)

	//admin routes
	adminGroup := r.Group("/admin")
	adminGroup.Use(
		middlewares.AuthMiddleware(),
		middlewares.AuthorizationRole("admin"),
	)

	adminGroup.POST("/fruits", handler.CreateFruits)
	adminGroup.PUT("/fruits/:id", handler.UpdateFruit)
	adminGroup.DELETE("/fruits/:id", handler.DeleteFruit)
}
