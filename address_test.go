package main

import (
	"testing"
)

func TestAddress(t *testing.T) {

	addr, err := NewAddress(1, 0)
	if err != nil {
		t.Error(err.Error())
	}

	if !ValidateAddressIdentifier(addr.Identifier) {
		t.Error("Invalid address identifier", addr.Identifier)
	}

	valid, err := ValidateAddress(addr.Identifier)
	if err != nil {
		t.Error(err.Error())
	}

	if valid != true {
		t.Error("Invalid checksum")
	}
}
