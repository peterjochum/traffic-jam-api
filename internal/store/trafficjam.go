package store

import (
	"github.com/peterjochum/traffic-jam-api/internal/models"
)

// TrafficJamStore has methods to provide backend functions for the TrafficJam-API
type TrafficJamStore interface {
	AddTrafficJam(jam models.TrafficJam) error
	GetTrafficJam(id int64) (models.TrafficJam, error)
	UpdateTrafficJam(id int64, jam models.TrafficJam) error
	DeleteTrafficJam(id int64)
	ListTrafficJams() []models.TrafficJam
}
