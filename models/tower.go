package models

type Tower struct {
	Id      string  `json:"id"`
	UserId  string  `json:"user_id"`
	Faction int64   `json:"faction"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
