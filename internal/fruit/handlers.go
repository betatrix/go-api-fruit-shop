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

func (h *FruitHandler) CreateFruits(c *gin.Context) {
	var data []FruitRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON:" + err.Error()})
		return
	}

	fruit, err := h.service.CreateFruits(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fruit)
}

func (h *FruitHandler) GetFruitbyID(c *gin.Context) {
	id := c.Param("id")

	fruitCreated, err := h.service.GetFruitbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruitCreated)
}

// TODO: cache with Redis
func (h *FruitHandler) GetAllFruits(c *gin.Context) {
	fruits, err := h.service.GetAllFruits()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func (h *FruitHandler) UpdateFruit(c *gin.Context) {
	var data FruitRequest
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

// TODO: soft delete
func (h *FruitHandler) DeleteFruit(c *gin.Context) {

}
