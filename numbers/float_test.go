package numbers

import "testing"

func TestRoundFloat(t *testing.T) {
	var testsTable = []struct {
		input    float64
		roundOn  float64
		places   int
		expected float64
	}{
		{12.1234, 0.5, 2, 12.12},
		{12.1234567, 0.5, 3, 12.123},
		{12.123123, 0.5, 4, 12.1231},
	}

	for _, tt := range testsTable {
		got := RoundFloat(tt.input, tt.roundOn, tt.places)
		if got != tt.expected {
			t.Errorf("Expected %f, got %f", tt.expected, got)
		}
	}
}
