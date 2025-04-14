package mutex

import "testing"

func TestRun(t *testing.T) {
	Run()
	if couter != 55 {
		t.Errorf("Invalid value to counter: %d. Expected: 55", couter)
	}
}
