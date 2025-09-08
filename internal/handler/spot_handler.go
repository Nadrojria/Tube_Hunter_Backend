package handler

import (
	"fmt"
	"net/http"
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, createdSpot)
}

func (handler *SpotHandler) UploadImage(context *gin.Context) {
	// Récupérer l'image depuis le formulaire
	file, err := context.FormFile("image")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Image not found"})
		return
	}

	// Créer un nom unique pour éviter les doublons
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// Sauvegarder le fichier (Gin le fait automatiquement)
	if err := context.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image"})
		return
	}

	// Retourner l'URL de l'image
	imageURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	context.JSON(http.StatusOK, gin.H{"photoUrl": imageURL})
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
