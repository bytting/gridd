package main

import (
	"testing"
)

func TestStore(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
