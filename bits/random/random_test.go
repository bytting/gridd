package random

import (
	"testing"
)

func TestRandom(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
