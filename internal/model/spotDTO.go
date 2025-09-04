package model

type SpotDTO struct {
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

type CreateSpotRequest struct {
	PhotoURL    string `json:"photoUrl" binding:"required"`
	Name        string `json:"name" binding:"required"`
	City        string `json:"city" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Difficulty  int    `json:"difficulty" binding:"required,min=1,max=5"`
	SurfBreaks  string `json:"surfBreaks" binding:"required"`
	SeasonStart string `json:"seasonStart" binding:"required"`
	SeasonEnd   string `json:"seasonEnd" binding:"required"`
}
