package proto

import (
//"bytes"
//"errors"
)

type Peer struct {
	IP        string
	Port      uint16
	PublicKey []byte
	Version   byte
}

func NewPeer() *Peer {

	return new(Peer)
}

func NewPeerFrom(ip string, port uint16) *Peer {

	peer := new(Peer)
	peer.IP = ip
	peer.Port = port

	return peer
}

func (p *Peer) Serialize() ([]byte, error) {
	return nil, nil
}

func (p *Peer) Deserialize(packet []byte) error {
	return nil
}
