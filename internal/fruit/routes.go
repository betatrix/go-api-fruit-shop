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
	r.GET("/fruits/:id", handler.GetFruitbyID)
	r.GET("/fruits", handler.GetAllFruits)
	r.PUT("/fruits/:id", handler.UpdateFruit)
	// r.DELETE("/fruits/{id}")

}
