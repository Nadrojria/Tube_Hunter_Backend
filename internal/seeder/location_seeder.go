package seeder

import (
    "database/sql"
    "log"
    "tubeHunter/internal/model"
)

var Locations = []model.Location{
    {
        ID: 1,
        Country: "France",
        City: "La Teste-de-Buch",
        Lat: 4,
        Long: 8,
    },
    {
        ID: 2,
        Country: "France",
        City: "Biscarosse",
        Lat: 4,
        Long: 9,
    },
}

func SeedLocations(db *sql.DB) {
    for _, s := range Locations {
        _, err := db.Exec(
            "INSERT OR IGNORE INTO locations (id, country, city, lat, long) VALUES (?, ?, ?, ?, ?)",
            s.ID, s.Country, s.City, s.Lat, s.Long,
        )
        if err != nil {
            log.Printf("⚠️ Erreur seed location %s: %v", s.City, err)
        }
    }
    log.Println("✅ Seeder executed")
}