package model

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
