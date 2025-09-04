package main

import (
	"tubeHunter/internal/handler"
	"tubeHunter/internal/repository"
	"tubeHunter/internal/seeder"
	"tubeHunter/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	db := pkg.InitDB("cmd/tubeHunter.db")

	seeder.SeedSpots(db)

	repo := &repository.SpotRepository{DB: db}
	handler := &handler.SpotHandler{Repo: repo}

	router := gin.Default()
	router.GET("/api/spots", handler.GetSpots)
	router.POST("/api/spots", handler.CreateSpot)
	router.GET("/api/spots/:id", handler.GetSpotByID)

	router.Run(":8080")
}
