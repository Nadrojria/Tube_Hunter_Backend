package handler

import (
	"net/http"
	"strings"
	"tubeHunter/internal/model"
	"tubeHunter/internal/repository"

	"github.com/gin-gonic/gin"
)

type SpotHandler struct {
	Repo *repository.SpotRepository
}

func (handler *SpotHandler) GetSpots(context *gin.Context) {
	spots, err := handler.Repo.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, spots)
}

func (handler *SpotHandler) CreateSpot(context *gin.Context) {
	var request model.Spot
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSpot, err := handler.Repo.Create(request)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			context.JSON(http.StatusConflict, gin.H{
				"error": "❌ Spot already exists",
			})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "✅ Spot added successfully",
		"data":    createdSpot,
	})
}
