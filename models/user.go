package models

type User struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	UberOpenId      string  `json:"uber_open_id"`
	Faction         int64   `json:"faction"`
	TotalScore      int64   `json:"total_score"`
	Score           int64   `json:"score"`
	LocationBearing int64   `json:"location_bearing"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	Hp              int64   `json:"hp"`
	Mana            int64   `json:"mana"`
}
