package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	// repo := NewUserRepository(db)
	// service := NewUserService(repo)
	// handler := NewUserHandler(service)

	// route.POST("/auth/login")
	// route.POST("/users")
	// route.GET("/users")
}
