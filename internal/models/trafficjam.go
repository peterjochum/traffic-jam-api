// Package models contains all data models for the application
package models

// TrafficJam stores information of a traffic problem at a specific location
type TrafficJam struct {
	ID                int64   `json:"id"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	DurationInSeconds int32   `json:"durationInSeconds"`
}
