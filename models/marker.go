package models

type Marker struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarkerPostRequest struct {
	Markers []Marker `json:"markers"`
}

type MarkerDeleteRequest struct {
	ID int `json:"id"`
}

type MarkerPutRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
