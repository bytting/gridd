package crypt

import (
	"bytes"
	"testing"
)

func TestCrypt(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	key := []byte("example key 1234")
	plaintext := []byte("exampleplaintext")

	enc, _ := AESEncrypt(plaintext, key)
	dec, _ := AESDecrypt(enc, key)

	if !bytes.Equal(plaintext, dec) {
		t.Errorf("AES decrypted data does not match. Expected %s, got %s", plaintext, dec)
	}
}
