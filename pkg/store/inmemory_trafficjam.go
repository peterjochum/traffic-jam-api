package store

import (
	"errors"
	"fmt"

	"github.com/peterjochum/traffic-jam-api/pkg/models"
)

// InMemoryTrafficJamStore allows stores TrafficJams in a map without persistence
type InMemoryTrafficJamStore struct {
	jams map[int64]models.TrafficJam
}

// NewInMemoryTrafficJamStore creates and optionally seeds the store
func NewInMemoryTrafficJamStore() *InMemoryTrafficJamStore {
	s := InMemoryTrafficJamStore{}
	s.jams = make(map[int64]models.TrafficJam)
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
func (i InMemoryTrafficJamStore) GetTrafficJam(id int64) (*models.TrafficJam, error) {
	if val, ok := i.jams[id]; ok {
		return &val, nil
	}
	return nil, errors.New("object not found")
}

// UpdateTrafficJam replaces information on the jam with the id
func (i InMemoryTrafficJamStore) UpdateTrafficJam(id int64, jam models.TrafficJam) error {
	if _, ok := i.jams[id]; ok {
		i.jams[jam.ID] = jam
		return nil
	}
	text := fmt.Sprintf("cannot update non existing jam %d", jam.ID)
	return errors.New(text)
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

// Total returns the number of jams in the store
func (i InMemoryTrafficJamStore) Total() int64 {
	return int64(len(i.jams))
}
