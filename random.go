// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package random

import (
	"crypto/rand"
	"io"
)

/* Generates a block of random bits */
func Bytes(size int) ([]byte, error) {

	b := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		b = nil
	}

	return b, err
}

func Entropy192() ([]byte, error) {

	b, err := Bytes(24)
	if err != nil {
		return nil, err
	}

	return b, nil
}
