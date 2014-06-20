// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"strings"

	"gridd/enc"
)

type Address struct {
	Version, Privacy byte
	Identifier       string
	Key              *ecdsa.PrivateKey
}

func NewAddress(version, privacy byte) (*Address, error) {

	var err error
	addr := new(Address)
	addr.Key, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, errors.New("address.NewAddress: Error generating ecdsa keys: " + err.Error())
	}

	addr.Version = version
	addr.Privacy = privacy

	var ident bytes.Buffer
	ident.WriteByte(addr.Version)
	ident.WriteByte(addr.Privacy)
	ident.Write(addr.Key.PublicKey.X.Bytes())
	ident.Write(addr.Key.PublicKey.Y.Bytes())
	cs, err := Checksum(ident.Bytes(), 2)
	if err != nil {
		return nil, errors.New("address.NewAddress: Error generating checksum: " + err.Error())
	}
	ident.Write(cs)

	ident58, err := enc.EncodeBase58(ident.Bytes())
	if err != nil {
		return nil, errors.New("address.NewAddress: Error base58 encoding: " + err.Error())
	}
	addr.Identifier = "LM:" + ident58
	return addr, nil
}

func ValidateAddressIdentifier(identifier string) bool {

	if len(identifier) < 7 { // LM: + version + privacy + checksum
		return false
	}

	if !strings.HasPrefix(identifier, "LM:") {
		return false
	}

	return true
}

func ValidateAddress(identifier string) (bool, error) {

	if !ValidateAddressIdentifier(identifier) {
		return false, nil
	}

	raw, err := enc.DecodeBase58(identifier[3:])
	if err != nil {
		return false, errors.New("address.ValidateChecksum: Error base58 decoding: " + err.Error())
	}
	ident := raw[:len(raw)-2]
	cs1 := raw[len(raw)-2:]
	cs2, err := Checksum(ident, 2)
	if err != nil {
		return false, errors.New("address.ValidateChecksum: Error generating checksum: " + err.Error())
	}

	return bytes.Compare(cs1, cs2) == 0, nil
}
