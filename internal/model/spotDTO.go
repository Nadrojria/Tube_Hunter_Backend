package model

type SpotDTO struct {
	ID          int         `json:"id"`
	PhotoURL    string      `json:"photoUrl"`
	Name        string      `json:"name"`
	Location    LocationDTO `json:"location"`
	Difficulty  int         `json:"difficulty"`
	SurfBreaks  string      `json:"surfBreaks"`
	SeasonStart string      `json:"seasonStart"`
	SeasonEnd   string      `json:"seasonEnd"`
}
