package fruit

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFruitRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewFruitRepository(db)
	service := NewFruitService(repo)
	handler := NewFruitHandler(service)

	r.POST("/fruits", handler.CreateFruits)
	// r.GET("/fruit/:id", handler.GetFruit)
	// r.GET("/fruits/{id}")
	// r.PUT("/fruits/{id}")
	// r.DELETE("/fruits/{id}")

}
