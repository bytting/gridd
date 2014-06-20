// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package proto

import (
//"bytes"
//"errors"
)

type ConnectMessage struct {
	Timestamp int64
}

func NewConnectMessage() *ConnectMessage {

	return new(ConnectMessage)
}

func (cm *ConnectMessage) Serialize() ([]byte, error) {
	return nil, nil
}

func (cm *ConnectMessage) Deserialize(packet []byte) error {
	return nil
}
