package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	r.POST("/users", handler.CreateUser)
	// r.POST("/auth/login")
	// r.GET("/users")
}
