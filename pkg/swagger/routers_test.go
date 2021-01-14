package swagger

import (
	"fmt"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	testcases := []struct {
		routeName string
	}{
		{"Index"},
		{"AddTrafficJam"},
		{"DeleteTrafficJam"},
		{"GetAllTrafficJams"},
		{"GetTrafficJam"},
		{"PutTrafficJam"},
	}

	for _, tc := range testcases {
		testName := fmt.Sprintf("route %s exists", tc.routeName)
		t.Run(testName, func(t *testing.T) {
			if router.GetRoute(tc.routeName) == nil {
				t.Errorf("No %s route found", tc.routeName)
			}
		})
	}

}
