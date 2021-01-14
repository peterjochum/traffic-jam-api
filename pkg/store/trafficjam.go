package store

import (
	"github.com/peterjochum/traffic-jam-api/pkg/models"
)

// TrafficJamStore has methods to provide backend functions for the TrafficJam-API
type TrafficJamStore interface {
	AddTrafficJam(jam models.TrafficJam) error
	GetTrafficJam(id int64) (*models.TrafficJam, error)
	UpdateTrafficJam(id int64, jam models.TrafficJam) error
	DeleteTrafficJam(id int64)
	ListTrafficJams() []models.TrafficJam
	Total() int64
}

// SeedTrafficJamStore inserts a handful of test TrafficJam objects
// for unit testing and development
func SeedTrafficJamStore(tjs TrafficJamStore) {
	_ = tjs.AddTrafficJam(models.TrafficJam{ID: 1, Longitude: 1.12, Latitude: 2.13, DurationInSeconds: 15})
	_ = tjs.AddTrafficJam(models.TrafficJam{ID: 2, Longitude: 2.12, Latitude: 3.13, DurationInSeconds: 30})
	_ = tjs.AddTrafficJam(models.TrafficJam{ID: 3, Longitude: 3.12, Latitude: 4.13, DurationInSeconds: 45})
}
