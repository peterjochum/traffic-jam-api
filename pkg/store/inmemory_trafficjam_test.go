package store

import (
	"fmt"
	"testing"

	"github.com/peterjochum/traffic-jam-api/pkg/models"
)

func TestNewInMemoryTrafficJamStore_Empty(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()

	if len(tjs.jams) > 0 {
		t.Errorf("traffic jam store should be empty")
	}
}

func TestNewInMemoryTrafficJamStore_Seed(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)

	expectedJams := 3
	if len(tjs.jams) != expectedJams {
		t.Errorf("store data structur should have %d elements", expectedJams)
	}
}

func TestInMemoryTrafficJamStore_ListTrafficJams(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)

	jams := tjs.ListTrafficJams()

	expectedJams := 3
	if len(jams) != expectedJams {
		t.Errorf("Expected %d traffic jams", expectedJams)
	}
}

func TestInMemoryTrafficJamStore_GetTrafficJam_Success(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)
	var expectedJam int64 = 1
	tj, err := tjs.GetTrafficJam(expectedJam)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if tj.ID != expectedJam {
		t.Errorf("unexpected traffic jam %d returned - expected %d", tj.ID, expectedJam)
	}
}

func TestInMemoryTrafficJamStore_GetTrafficJam_NonExisting(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	tj, err := tjs.GetTrafficJam(1)
	if err == nil {
		t.Errorf("expected error, but call was success")
	}

	if tj != nil {
		t.Errorf("Traffic jam returned, how?")
	}
}

func TestInMemoryTrafficJamStore_AddTrafficJam(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	var jamID int64 = 1
	err := tjs.AddTrafficJam(models.TrafficJam{ID: jamID})
	if err != nil {
		t.Error(err)
	}

	expectedJamCount := 1
	if len(tjs.jams) != expectedJamCount {
		t.Errorf("expected the jam %d to have been added to the store", jamID)
	}
}

func TestInMemoryTrafficJamStore_AddTrafficJam_Existing(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)

	// New traffic jam with existing id
	tj := models.TrafficJam{ID: 1}

	err := tjs.AddTrafficJam(tj)
	if err == nil {
		t.Errorf("expected error: traffic jam exists")
	} else {
		expectedError := fmt.Sprintf("jam %d already exists", 1)
		if err.Error() != expectedError {
			t.Errorf("expected error \"%s\", but got \"%s\"",
				expectedError, err.Error())
		}
	}
}

func TestInMemoryTrafficJamStore_UpdateTrafficJam(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)

	tj := models.TrafficJam{ID: 1}
	err := tjs.UpdateTrafficJam(1, tj)
	if err != nil {
		t.Error(err)
	}
}

func TestInMemoryTrafficJamStore_UpdateTrafficJam_Existing(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	err := tjs.UpdateTrafficJam(1, models.TrafficJam{})
	if err == nil {
		t.Errorf("updating non-existing jam should lead to error")
	}
}

func TestInMemoryTrafficJamStore_DeleteTrafficJam(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)
	var testJamID int64 = 1
	tjs.DeleteTrafficJam(testJamID)

	expectedJamCount := 2
	if len(tjs.jams) != expectedJamCount {
		t.Errorf("expected %d jams to be left after deletion, got %d",
			expectedJamCount, len(tjs.jams))
	}

	jam, err := tjs.GetTrafficJam(testJamID)
	if err == nil {
		t.Errorf("getting deleted jam %d should result in error", testJamID)
	}

	if jam != nil {
		t.Errorf("no jam should be returned")
	}
}
