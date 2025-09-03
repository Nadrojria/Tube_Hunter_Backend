package model

type SpotDTO struct {
	ID          int         `json:"id"`
	PhotoURL    string      `json:"photo_url"`
	Name        string      `json:"name"`
	Location    LocationDTO `json:"location"`
	Difficulty  int         `json:"difficulty"`
	SurfBreaks  string      `json:"surf_breaks"`
	SeasonStart string      `json:"season_start"`
	SeasonEnd   string      `json:"season_end"`
}
