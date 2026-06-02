package handler

import (
	"net/http"
	"strings"

	"github.com/PradeepSahhu/NeuralTraderX/internal/models"
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
	"github.com/gin-gonic/gin"
)

type TempHandler struct {
	tempRepo *repository.TempRepository
}

type createTempRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewTempHandler(repo *repository.TempRepository) *TempHandler {
	return &TempHandler{
		tempRepo: repo,
	}
}

// Register Routes to Handler functions
func (h *TempHandler) RegisterRoutes(rg *gin.RouterGroup) {
	temp := rg.Group("/temps")

	temp.POST("/", h.CreateTempHandler)

}

// functions

func (h *TempHandler) CreateTempHandler(c *gin.Context) {
	var request createTempRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	if strings.TrimSpace(request.FirstName) == "" || strings.TrimSpace(request.LastName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "firstName and lastName are required"})
		return
	}

	temp := models.Temp{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	if err := h.tempRepo.CreateTemp(&temp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unsuccessfully to inserted the record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully inserted the record"})
}
