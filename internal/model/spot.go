package model

import "time"

type Spot struct {
    ID          int       `json:"id"`
    PhotoURL    string    `json:"photo_url"`
    Name        string    `json:"name"`
    LocationID  int       `json:"location_id"`
    Difficulty  int       `json:"difficulty"`
    SurfBreaks  string    `json:"surf_breaks"`
    SeasonStart time.Time `json:"season_start"`
    SeasonEnd   time.Time `json:"season_end"`
}