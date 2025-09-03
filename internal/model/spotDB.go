package model

type SpotDB struct {
	ID          int    `json:"id"`
	PhotoURL    string `json:"photoUrl"`
	Name        string `json:"name"`
	LocationID  int    `json:"locationId"`
	Difficulty  int    `json:"difficulty"`
	SurfBreaks  string `json:"surfBreaks"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
}
