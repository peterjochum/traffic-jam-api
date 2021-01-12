package store

import (
	"errors"
	"fmt"

	"github.com/peterjochum/traffic-jam-api/internal/models"
)

// InMemoryTrafficJamStore allows stores TrafficJams in a map without persistence
type InMemoryTrafficJamStore struct {
	jams map[int64]models.TrafficJam
}

// NewInMemoryTrafficJamStore creates and optionally seeds the store
func NewInMemoryTrafficJamStore(seed bool) *InMemoryTrafficJamStore {
	s := InMemoryTrafficJamStore{}
	s.jams = make(map[int64]models.TrafficJam)
	if seed {
		s.seed()
	}
	return &s
}

// AddTrafficJam adds a new traffic jam to the map
func (i InMemoryTrafficJamStore) AddTrafficJam(jam models.TrafficJam) error {
	_, exist := i.jams[jam.ID]
	if exist {
		return fmt.Errorf("jam %d already exists", jam.ID)
	}
	i.jams[jam.ID] = jam
	return nil
}

// GetTrafficJam returns a traffic jam with the specified id
func (i InMemoryTrafficJamStore) GetTrafficJam(id int64) (models.TrafficJam, error) {
	val, ok := i.jams[id]
	if ok {
		return val, nil
	}
	return models.TrafficJam{}, errors.New("object not found")
}

// UpdateTrafficJam replaces information on the jam with the id
func (i InMemoryTrafficJamStore) UpdateTrafficJam(id int64, jam models.TrafficJam) error {
	_, exist := i.jams[id]
	if !exist {
		text := fmt.Sprintf("cannot update non existing jam %d", jam.ID)
		return errors.New(text)
	}
	i.jams[jam.ID] = jam
	return nil
}

// DeleteTrafficJam removes a traffic jam by id
func (i InMemoryTrafficJamStore) DeleteTrafficJam(id int64) {
	delete(i.jams, id)
}

// ListTrafficJams returns a list of all available traffic jams
func (i InMemoryTrafficJamStore) ListTrafficJams() []models.TrafficJam {
	var jamList []models.TrafficJam
	for _, v := range i.jams {
		jamList = append(jamList, v)
	}
	return jamList
}

// seed inserts a handful of test TrafficJam objects for unit testing and development
func (i InMemoryTrafficJamStore) seed() {
	_ = i.AddTrafficJam(models.TrafficJam{ID: 1, Longitude: 1.12, Latitude: 2.13, DurationInSeconds: 15})
	_ = i.AddTrafficJam(models.TrafficJam{ID: 2, Longitude: 2.12, Latitude: 3.13, DurationInSeconds: 30})
	_ = i.AddTrafficJam(models.TrafficJam{ID: 3, Longitude: 3.12, Latitude: 4.13, DurationInSeconds: 45})
}
