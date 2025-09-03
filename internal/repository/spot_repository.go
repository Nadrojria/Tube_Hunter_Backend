package repository

import (
	"database/sql"
	"tubeHunter/internal/model"
)

type SpotRepository struct {
	DB *sql.DB
}

func (repo *SpotRepository) GetAll() ([]model.SpotDTO, error) {
	query := `
        SELECT s.id, s.photo_url, s.name, s.difficulty, s.surf_breaks, s.season_start, s.season_end,
               l.country, l.city, l.lat, l.long
        FROM spots s 
        JOIN locations l ON s.location_id = l.id
    `

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []model.SpotDTO
	for rows.Next() {
		var s model.SpotDTO
		var l model.LocationDTO

		err := rows.Scan(
			&s.ID,
			&s.PhotoURL,
			&s.Name,
			&s.Difficulty,
			&s.SurfBreaks,
			&s.SeasonStart,
			&s.SeasonEnd,
			&l.Country,
			&l.City,
			&l.Lat,
			&l.Long,
		)

		if err != nil {
			return nil, err
		}

		s.Location = l
		spots = append(spots, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

func (repo *SpotRepository) GetByID(spotID int) (*model.SpotDTO, error) {
	query := `
        SELECT 
            s.id, s.photo_url, s.name, s.difficulty, 
            s.surf_breaks, s.season_start, s.season_end,
            l.country, l.city, l.lat, l.long
        FROM spots s
        JOIN locations l ON s.location_id = l.id
        WHERE s.id = ?`

	var spot model.SpotDTO
	var location model.LocationDTO

	row := repo.DB.QueryRow(query, spotID)

	err := row.Scan(
		&spot.ID,
		&spot.PhotoURL,
		&spot.Name,
		&spot.Difficulty,
		&spot.SurfBreaks,
		&spot.SeasonStart,
		&spot.SeasonEnd,
		&location.Country,
		&location.City,
		&location.Lat,
		&location.Long,
	)

	if err != nil {
		return nil, err
	}

	spot.Location = location

	return &spot, nil
}

func (repo *SpotRepository) Create(spot model.SpotDTO) error {
	query := `INSERT INTO spots (
		photo_url, 
		name, 
		location_id, 
		difficulty, 
		surf_breaks, 
		season_start, 
		season_end) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
		`
	_, err := repo.DB.Exec(query,
		spot.PhotoURL,
		spot.Name,

		spot.Difficulty,
		spot.SurfBreaks,
		spot.SeasonStart,
		spot.SeasonEnd)
	return err
}
