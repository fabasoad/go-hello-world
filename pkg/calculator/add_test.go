package calculator

import "testing"

func TestAdd(t *testing.T) {
	x := 3
	y := 5
	result := Add(x, y)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}
