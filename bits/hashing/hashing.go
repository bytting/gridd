// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package hashing

import (
	"crypto/sha256"

	"gridd/bits/hashing/ripemd160"
)

func SHA256(p []byte) []byte {

	sha := sha256.New()
	sha.Write(p)

	return sha.Sum(nil)
}

func RIPEMD160(p []byte) []byte {

	ripe := ripemd160.New()
	ripe.Write(p)

	return ripe.Sum(nil)
}

func Hash(p []byte) []byte {

	return RIPEMD160(SHA256(p))
}
