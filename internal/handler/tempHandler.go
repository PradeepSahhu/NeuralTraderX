package handler

import (
	"net/http"

	"github.com/PradeepSahhu/NeuralTraderX/internal/models"
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
	"github.com/gin-gonic/gin"
)

type TempHandler struct {
	tempRepo *repository.TempRepository
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

	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")

	temp := models.Temp{
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: "2026-01-01",
	}

	if err := h.tempRepo.CreateTemp(&temp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unsuccessfully to inserted the record"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully inserted the record"})
}
