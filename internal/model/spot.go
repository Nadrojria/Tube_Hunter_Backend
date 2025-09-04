package model

type Spot struct {
	ID          int    `json:"id"`
	PhotoURL    string `json:"photoUrl"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Difficulty  int    `json:"difficulty"`
	SurfBreaks  string `json:"surfBreaks"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
}
