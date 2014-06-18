// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

/*
AES decrypts data. This is done in-place, this function has a side effect.
*/
func AESDecrypt(data []byte, key []byte) ([]byte, error) {

	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("Data should be a multiple of 128 bits")
	}

	if len(data) < 32 {
		return nil, errors.New("iv + data should be at least 256 bits")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := data[:16]
	data = data[16:]

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(data, data)

	return data, nil
}

/*
AES encrypts data.
*/
func AESEncrypt(data []byte, key []byte) ([]byte, error) {

	if len(data)%16 != 0 {
		return nil, errors.New("Data should be a multiple of 128 bits")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(data)+16)

	iv := ciphertext[:16]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[16:], data)

	return ciphertext, nil
}
