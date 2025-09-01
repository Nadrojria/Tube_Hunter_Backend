package main

import (
    "github.com/gin-gonic/gin"
    "tubeHunter/internal/handler"
    "tubeHunter/internal/repository"
    "tubeHunter/internal/seeder"
    "tubeHunter/pkg"
)

func main() {
    db := pkg.InitDB("./tubeHunter.db")
    
    seeder.SeedSpots(db)
    seeder.SeedLocations(db)

    repo := &repository.SpotRepository{DB: db}
    handler := &handler.SpotHandler{Repo: repo}

    router := gin.Default()
    router.GET("/spots", handler.GetSpots)
    router.POST("/spots", handler.CreateSpot)

    router.Run(":8080")
}
