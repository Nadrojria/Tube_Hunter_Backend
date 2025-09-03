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

type CreateSpotRequest struct {
	PhotoURL    string `json:"photoUrl" binding:"required"`
	Name        string `json:"name" binding:"required"`
	LocationID  int    `json:"locationId" binding:"required"`
	Difficulty  int    `json:"difficulty" binding:"required,min=1,max=5"`
	SurfBreaks  string `json:"surfBreaks"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
}
