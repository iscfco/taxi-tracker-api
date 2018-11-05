package model

type Vehicle struct {
	Id        string  `json:"id"`
	Make      string  `json:"make"`
	Model     string  `json:"model"`
	Year      int16   `json:"year"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
