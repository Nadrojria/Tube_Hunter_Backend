package seeder

import (
    "database/sql"
    "log"
    "tubeHunter/internal/model"
    "time"
)

var Spots = []model.Spot{
    {
        ID:          1,
        PhotoURL:    "https://res.cloudinary.com/manawa/image/private/f_auto,c_limit,w_3840,q_auto/aykvlohikeutpdcp720o",
        Name:        "Le petit Nice",
        LocationID:  1,
        Difficulty:  4,
        SurfBreaks:  "Reef Break",
        SeasonStart: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC),
        SeasonEnd:   time.Date(2025, time.September, 30, 0, 0, 0, 0, time.UTC),
    },
    {
        ID:          2,
        PhotoURL:    "https://res.cloudinary.com/manawa/image/private/f_auto,c_limit,w_3840,q_auto/aykvlohikeutpdcp720o",
        Name:        "Nazare",
        LocationID:  2,
        Difficulty:  5,
        SurfBreaks:  "Point Break",
        SeasonStart: time.Date(2025, time.October, 1, 0, 0, 0, 0, time.UTC),
        SeasonEnd:   time.Date(2026, time.February, 28, 0, 0, 0, 0, time.UTC),
    },
}

func SeedSpots(db *sql.DB) {
    for _, s := range Spots {
        _, err := db.Exec(
            "INSERT OR IGNORE INTO spots (id, photo_url, name, location_id, difficulty, surf_breaks, season_start, season_end) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
            s.ID, s.PhotoURL, s.Name, s.LocationID, s.Difficulty, s.SurfBreaks, s.SeasonStart, s.SeasonEnd,
        )
        if err != nil {
            log.Printf("⚠️ Erreur seed spot %s: %v", s.Name, err)
        }
    }
    log.Println("✅ Seeder executed")
}