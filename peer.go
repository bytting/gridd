// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package main

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
