package model

type LocationDB struct {
	ID      int     `json:"id"`
	Country string  `json:"country"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
}
