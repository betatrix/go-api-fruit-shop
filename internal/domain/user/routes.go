package user

import (
	"github.com/betatrix/go-api-fruit-shop/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	//public routes
	publicGroup := r.Group("/")
	publicGroup.POST("/auth/login", handler.LoginHandler)

	//admin routes
	adminGroup := r.Group("/admin")
	adminGroup.Use(
		middlewares.AuthMiddleware(),
		middlewares.AuthorizationRole("admin"),
	)

	adminGroup.POST("/users", handler.CreateUser)
	adminGroup.GET("/users/:id", handler.GetUserByID)
	adminGroup.GET("/users/", handler.GetAllUsers)
}
