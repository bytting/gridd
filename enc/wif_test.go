package enc

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestWIF(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	keys1, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		t.Error(err.Error())
	}

	wif, err := EncodeWif(keys1)
	if err != nil {
		t.Error(err.Error())
	}

	keys2, err := DecodeWif(wif)
	if err != nil {
		t.Error(err.Error())
	}

	if bytes.Compare(keys1.D.Bytes(), keys2.D.Bytes()) != 0 {
		t.Error("Private keys are different. Expected %x, got %x\n", keys1.D.Bytes(), keys2.D.Bytes())
	}

	if bytes.Compare(keys1.PublicKey.X.Bytes(), keys2.PublicKey.X.Bytes()) != 0 {
		t.Error("Public point X are different. Expected %x, got %x\n", keys1.PublicKey.X.Bytes(), keys2.PublicKey.X.Bytes())
	}

	if bytes.Compare(keys1.PublicKey.Y.Bytes(), keys2.PublicKey.Y.Bytes()) != 0 {
		t.Error("Public point Y are different. Expected %x, got %x\n", keys1.PublicKey.Y.Bytes(), keys2.PublicKey.Y.Bytes())
	}

	if !keys2.PublicKey.Curve.IsOnCurve(keys2.PublicKey.X, keys2.PublicKey.Y) {
		t.Error("Public point is not on curve\n")
	}

	ok, err := ValidateWif(wif)
	if err != nil {
		t.Error(err.Error())
	}

	if !ok {
		t.Error("Invalid checksum")
	}
}
