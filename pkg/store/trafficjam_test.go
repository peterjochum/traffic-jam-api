package store

import "testing"

func TestSeedTrafficJamStore(t *testing.T) {
	tjs := NewInMemoryTrafficJamStore()
	SeedTrafficJamStore(tjs)

	if len(tjs.jams) != 3 {
		t.Errorf("expected %d jams in the store after seeding, got %d",
			3, len(tjs.jams))
	}
}
