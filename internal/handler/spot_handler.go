package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "tubeHunter/internal/model"
    "tubeHunter/internal/repository"
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
    var spot model.Spot
    if err := context.ShouldBindJSON(&spot); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := handler.Repo.Create(spot); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusCreated, spot)
}
