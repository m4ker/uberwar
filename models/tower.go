package models

type Tower struct {
	Id     string  `json:"id"`
	UserId string  `json:"user_id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}
