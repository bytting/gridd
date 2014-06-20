// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package main

import (
	"crypto/rand"
	"errors"
	"io"
)

func Checksum(p []byte, n int) ([]byte, error) {

	if n > len(p) {
		return nil, errors.New("bits.Checksum: checksum size is bigger than digest length")
	}

	npad := n - (len(p) % n)
	check := make([]byte, n)
	buf := make([]byte, len(p)+npad)
	copy(buf, p)

	for i := 0; i < len(buf); i += n {
		for j := 0; j < n; j++ {
			check[j] ^= buf[i+j]
		}
	}

	return check, nil
}

/* Generates a block of random bits */
func RandomBytes(size int) ([]byte, error) {

	b := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		b = nil
	}

	return b, err
}

func Entropy192() ([]byte, error) {

	b, err := RandomBytes(24)
	if err != nil {
		return nil, err
	}

	return b, nil
}
