package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var data UserDTO

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON:" + err.Error()})
		return
	}

	userCreated, err := h.service.CreateUser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userCreated)
}
