// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package enc

import (
	"bytes"
	"errors"
	"math/big"
	"strings"
)

const (
	alphabet58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func EncodeBase58(b []byte) (string, error) {

	if len(b) < 1 {
		return "", errors.New("base58.Encode: Byte slice is too short")
	}
	zero := big.NewInt(0)
	val := big.NewInt(0)
	val.SetBytes(b)

	var buffer bytes.Buffer

	if val.Cmp(zero) == 0 {
		buffer.WriteByte(alphabet58[0])
		return buffer.String(), nil
	}

	n := val
	r := big.NewInt(0)
	base := big.NewInt(58)

	for n.Cmp(zero) > 0 {
		r.Mod(n, base)
		n.Div(n, base) // FIXME: Use DivMod
		buffer.WriteByte(alphabet58[r.Uint64()])
	}

	length := len(b)
	for i := 0; i < length && b[i] == 0; i++ {
		buffer.WriteByte(alphabet58[0])
	}

	length = len(buffer.Bytes())
	for i := 0; i < length/2; i++ {
		buffer.Bytes()[i], buffer.Bytes()[length-1-i] = buffer.Bytes()[length-1-i], buffer.Bytes()[i]
	}

	return buffer.String(), nil
}

func DecodeBase58(encoded string) ([]byte, error) {

	bn := big.NewInt(0)
	base := big.NewInt(58)
	tmp := big.NewInt(0)

	for i := 0; i < len(encoded); i++ {
		pos := strings.IndexRune(alphabet58, rune(encoded[i]))
		if pos == -1 {
			return nil, errors.New("base58.Decode: Character not present in base58")
		}
		tmp.SetUint64(uint64(pos))

		bn.Mul(bn, base)
		bn.Add(bn, tmp)
	}

	var buf bytes.Buffer
	length := len(encoded)
	for i := 0; i < length && encoded[i] == alphabet58[0]; i++ {
		buf.WriteByte(0)
	}
	buf.Write(bn.Bytes())

	return buf.Bytes(), nil
}
