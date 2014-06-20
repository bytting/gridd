// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package wif

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"errors"
	"math/big"

	"gridd/bits/encoding/base58"
)

func Encode(keys *ecdsa.PrivateKey) (string, error) {

	var extended bytes.Buffer
	extended.WriteByte(byte(0x80))
	extended.Write(keys.D.Bytes())
	sha1, sha2 := sha256.New(), sha256.New()
	sha1.Write(extended.Bytes())
	sha2.Write(sha1.Sum(nil))
	checksum := sha2.Sum(nil)[:4]
	extended.Write(checksum)
	encoded, err := base58.Encode(extended.Bytes())
	if err != nil {
		return "", err
	}
	return encoded, nil
}

func Decode(wif string) (*ecdsa.PrivateKey, error) {

	if len(wif) < 6 {
		return nil, errors.New("wif.Decode: wif is too short")
	}

	extended, err := base58.Decode(wif)
	if err != nil {
		return nil, err
	}

	decoded := extended[1 : len(extended)-4]
	keys := new(ecdsa.PrivateKey)
	keys.D = new(big.Int).SetBytes(decoded)
	keys.PublicKey.Curve = elliptic.P384()
	for keys.PublicKey.X == nil {
		keys.PublicKey.X, keys.PublicKey.Y = keys.PublicKey.Curve.ScalarBaseMult(decoded)
	}

	if !keys.Curve.IsOnCurve(keys.PublicKey.X, keys.PublicKey.Y) {
		return nil, errors.New("wif.Decode: Point is not on curve")
	}
	return keys, nil
}

func ValidateChecksum(wif string) (bool, error) {

	if len(wif) < 6 {
		return false, errors.New("wif.Validate: wif is too short")
	}

	extended, err := base58.Decode(wif)
	if err != nil {
		return false, err
	}

	cs1 := extended[len(extended)-4:]
	sha1, sha2 := sha256.New(), sha256.New()
	sha1.Write(extended[:len(extended)-4])
	sha2.Write(sha1.Sum(nil))
	cs2 := sha2.Sum(nil)[:4]
	return bytes.Compare(cs1, cs2) == 0, nil
}
