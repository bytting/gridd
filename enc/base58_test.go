package enc

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func TestBase58(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var buf bytes.Buffer
	buf.WriteString("This is a string")

	encoded, err := Encode(buf.Bytes())
	if err != nil {
		t.Error(err.Error())
	}

	decoded, err := Decode(encoded)
	if err != nil {
		t.Error(err.Error())
	}

	if string(decoded) != buf.String() {
		t.Errorf("Decoded base58 does not match. Expected %s, got %s", buf.String(), string(decoded))
	}

	b := []byte{0, 0, 0, 1, 2, 3}
	encoded, _ = Encode(b)
	decoded, _ = Decode(encoded)
	if bytes.Compare(b, decoded) != 0 {
		t.Error("Decoded base58 with zeros does not match")
	}

	for j := 0; j < 10; j++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b2 := make([]byte, 253)
		for i := 0; i < len(b2); i++ {
			b2[i] = byte(r.Int())
		}
		encoded, _ = Encode(b2)
		decoded, _ = Decode(encoded)
		if bytes.Compare(b2, decoded) != 0 {
			t.Error("Decoded base58 with rand does not match")
		}
	}
}
