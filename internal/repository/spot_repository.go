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
        SELECT s.id, s.photo_url, s.name, s.city, s.country, s.difficulty, s.surf_breaks, s.season_start, s.season_end
		FROM spots s
`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	spots := make([]model.Spot, 0)
	for rows.Next() {
		var s model.Spot

		err := rows.Scan(
			&s.ID,
			&s.PhotoURL,
			&s.Name,
			&s.City,
			&s.Country,
			&s.Difficulty,
			&s.SurfBreaks,
			&s.SeasonStart,
			&s.SeasonEnd,
		)

		if err != nil {
			return nil, err
		}

		spots = append(spots, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

func (repo *SpotRepository) Create(request model.Spot) (*model.Spot, error) {

	query := `INSERT INTO spots (photo_url, name, city, country, difficulty, surf_breaks, season_start, season_end) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := repo.DB.Exec(query,
		request.PhotoURL,
		request.Name,
		request.City,
		request.Country,
		request.Difficulty,
		request.SurfBreaks,
		request.SeasonStart,
		request.SeasonEnd)

	if err != nil {
		return nil, err
	}

	spotID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	query = `
        SELECT s.id, s.photo_url, s.name, s.city, s.country, s.difficulty, 
		s.surf_breaks, s.season_start, s.season_end     
        FROM spots s 
        WHERE s.id = ?
    `

	var s model.Spot

	err = repo.DB.QueryRow(query, spotID).Scan(
		&s.ID,
		&s.PhotoURL,
		&s.Name,
		&s.City,
		&s.Country,
		&s.Difficulty,
		&s.SurfBreaks,
		&s.SeasonStart,
		&s.SeasonEnd,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
