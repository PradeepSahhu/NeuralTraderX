package handler

import (
	"net/http"

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

func (h *TempHandler) CreateTempHandler(c *gin.Context) {

	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")

	if err := h.tempRepo.CreateTemp(firstName, lastName); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully inserted the record"})
	}
}
