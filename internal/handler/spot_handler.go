package handler

import (
	"net/http"
	// "strconv"
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, createdSpot)
}

// func (handler *SpotHandler) GetSpotByID(context *gin.Context) {
// 	idStr := context.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
// 		return
// 	}

// 	spot, err := handler.Repo.GetByID(id)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if spot == nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "spot not found"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, spot)
// }
