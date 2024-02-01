package integers

import "testing"

func TestAdde(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}
