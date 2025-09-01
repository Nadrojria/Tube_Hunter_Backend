package repository

import (
	"database/sql"
	"tubeHunter/internal/model"
)

type SpotRepository struct {
	DB *sql.DB
}

func (repo *SpotRepository) GetAll() ([]model.Spot, error) {
	rows, err := repo.DB.Query("SELECT id, photo_url, name, location_id, difficulty, surf_breaks FROM spots")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []model.Spot
	for rows.Next() {
		var s model.Spot
		rows.Scan(&s.ID, &s.PhotoURL, &s.Name, &s.LocationID, &s.Difficulty, &s.SurfBreaks)
		spots = append(spots, s)
	}
	return spots, nil
}

func (repo *SpotRepository) Create(spot model.Spot) error {
	_, err := repo.DB.Exec("INSERT INTO spots (photo_url, name, location_id, difficulty, surf_breaks, season_start, season_end) VALUES (?, ?, ?, ?, ?, ?, ?)",
		spot.PhotoURL, spot.Name, spot.LocationID, spot.Difficulty, spot.SurfBreaks, spot.SeasonStart, spot.SeasonEnd)
	return err
}

func (repo *SpotRepository) GetByID(id int) (*model.Spot, error) {
	row := repo.DB.QueryRow(`
        SELECT id, photo_url, name, location_id, difficulty, surf_breaks, season_start, season_end
        FROM spots
        WHERE id = ?
    `, id)

	var s model.Spot
	err := row.Scan(&s.ID, &s.PhotoURL, &s.Name, &s.LocationID, &s.Difficulty, &s.SurfBreaks, &s.SeasonStart, &s.SeasonEnd)
	if err == sql.ErrNoRows {
		return nil, nil // pas trouvé
	} else if err != nil {
		return nil, err
	}

	return &s, nil
}
