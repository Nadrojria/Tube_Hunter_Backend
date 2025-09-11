package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
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

func (handler *SpotHandler) UploadImage(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Image not found"})
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// Save file (Gin does it automatically)
	if err := context.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image"})
		return
	}

	// URL construction from request(BASE_URL)
	scheme := "http"
	if context.Request.TLS != nil {
		scheme = "https"
	}

	// Automatically use IP/host where the request came from
	host := context.Request.Host

	imageURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, host, filename)
	context.JSON(http.StatusOK, gin.H{"photoUrl": imageURL})
}
