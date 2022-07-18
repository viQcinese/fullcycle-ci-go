package main
import "testing"

func TestAdd(t *testing.T) {
	total := Add(15, 15)

	if total != 30 {
		t.Errorf("Invalid result. Result: %d. Expected: %d", total, 30)
	}
}