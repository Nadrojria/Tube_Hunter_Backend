package seeder

import (
	"database/sql"
	"log"
	"tubeHunter/internal/model"
)

var Locations = []model.Location{
	{ID: 1, Country: "France", City: "Nice", Lat: 43.7, Long: 7.3},
	{ID: 2, Country: "Portugal", City: "Nazaré", Lat: 39.6, Long: -9.1},
	{ID: 3, Country: "France", City: "Hossegor", Lat: 43.66, Long: -1.44},
	{ID: 4, Country: "Hawaii", City: "Oahu", Lat: 21.66, Long: -158.05},
	{ID: 5, Country: "South Africa", City: "Jeffreys Bay", Lat: -34.05, Long: 24.91},
	{ID: 6, Country: "USA", City: "Mavericks", Lat: 37.49, Long: -122.50},
	{ID: 7, Country: "Tahiti", City: "Teahupo’o", Lat: -17.83, Long: -149.27},
	{ID: 8, Country: "Hawaii", City: "Banzai Pipeline", Lat: 21.66, Long: -158.05},
	{ID: 9, Country: "Indonesia", City: "Uluwatu", Lat: -8.83, Long: 115.09},
	{ID: 10, Country: "Australia", City: "Bondi Beach", Lat: -33.89, Long: 151.28},
	{ID: 11, Country: "Australia", City: "Snapper Rocks", Lat: -28.17, Long: 153.55},
	{ID: 12, Country: "Hawaii", City: "Waimea Bay", Lat: 21.64, Long: -158.06},
	{ID: 13, Country: "New Zealand", City: "Raglan", Lat: -37.80, Long: 174.87},
	{ID: 14, Country: "Tahiti", City: "Teahupo’o Left", Lat: -17.83, Long: -149.27},
	{ID: 15, Country: "Mexico", City: "Puerto Escondido", Lat: 15.87, Long: -97.09},
	{ID: 16, Country: "UK", City: "Fistral Beach", Lat: 50.42, Long: -5.08},
	{ID: 17, Country: "Australia", City: "Bells Beach", Lat: -38.37, Long: 144.28},
	{ID: 18, Country: "Indonesia", City: "Canggu", Lat: -8.65, Long: 115.13},
	{ID: 19, Country: "Puerto Rico", City: "Rincon", Lat: 18.34, Long: -67.25},
	{ID: 20, Country: "El Salvador", City: "El Sunzal", Lat: 13.31, Long: -89.33},
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
