package repository

import (
	"database/sql"
	"tubeHunter/internal/model"
)

type SpotRepository struct {
	DB *sql.DB
}

func (repo *SpotRepository) GetAll() ([]model.Spot, error) {
	query := `
        SELECT s.id, s.photo_url, s.name, s.location_id, s.difficulty, s.surf_breaks, s.season_start, s.season_end,
               l.id, l.country, l.city, l.lat, l.long
        FROM spots s 
        INNER JOIN locations l ON s.location_id = l.id
    `

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []model.Spot
	for rows.Next() {
		var s model.Spot
		var l model.Location

		err := rows.Scan(
			&s.ID, &s.PhotoURL, &s.Name, &s.LocationID, &s.Difficulty, &s.SurfBreaks, &s.SeasonStart, &s.SeasonEnd,
			&l.ID, &l.Country, &l.City, &l.Lat, &l.Long,
		)
		if err != nil {
			return nil, err
		}

		s.Location = l
		spots = append(spots, s)
	}

	return spots, nil
}

func (repo *SpotRepository) Create(spot model.Spot) error {
	_, err := repo.DB.Exec("INSERT INTO spots (photo_url, name, location_id, difficulty, surf_breaks, season_start, season_end) VALUES (?, ?, ?, ?, ?, ?, ?)",
		spot.PhotoURL,
		spot.Name,
		spot.LocationID,
		spot.Difficulty,
		spot.SurfBreaks,
		spot.SeasonStart,
		spot.SeasonEnd)
	return err
}

func (repo *SpotRepository) GetByID(id int) (*model.Spot, error) {
	query := `
        SELECT s.id, s.photo_url, s.name, s.location_id, s.difficulty, s.surf_breaks, s.season_start, s.season_end,
               l.id, l.country, l.city, l.lat, l.long
        FROM spots s 
        INNER JOIN locations l ON s.location_id = l.id
        WHERE s.id = ?
    `

	row := repo.DB.QueryRow(query, id)

	var s model.Spot
	var l model.Location

	err := row.Scan(
		&s.ID, &s.PhotoURL, &s.Name, &s.LocationID, &s.Difficulty, &s.SurfBreaks, &s.SeasonStart, &s.SeasonEnd,
		&l.ID, &l.Country, &l.City, &l.Lat, &l.Long,
	)

	if err == sql.ErrNoRows {
		return nil, nil //not found
	} else if err != nil {
		return nil, err
	}

	s.Location = l
	return &s, nil
}
