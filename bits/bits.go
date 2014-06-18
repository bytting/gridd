// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package bits

import (
	"errors"
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
