package fruit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FruitHandler struct {
	service *FruitService
}

func NewFruitHandler(service *FruitService) *FruitHandler {
	return &FruitHandler{service: service}
}

// TODO: Only admin
func (h *FruitHandler) CreateFruits(c *gin.Context) {
	var data FruitsDTO

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON:" + err.Error()})
		return
	}

	fruitCreated, err := h.service.CreateFruits(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fruitCreated)
}

// TODO: Admin and user
func (h *FruitHandler) GetFruitbyID(c *gin.Context) {
	id := c.Param("id")

	fruitCreated, err := h.service.GetFruitbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitCreated)
}

// TODO: Admin and user
// TODO: cache with Redis
func (h *FruitHandler) GetAllFruits(c *gin.Context) {
	fruits, err := h.service.GetAllFruits()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

// TODO: Only admin
func (h *FruitHandler) UpdateFruit(c *gin.Context) {
	var data FruitDTO
	id := c.Param("id")

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON:" + err.Error()})
		return
	}

	fruitUpdated, err := h.service.UpdateFruit(id, data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fruitUpdated)
}

// TODO: Only admin
func (h *FruitHandler) DeleteFruit(c *gin.Context) {
	id := c.Param("id")

	successMsg, err := h.service.DeleteFruit(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, successMsg)
}
