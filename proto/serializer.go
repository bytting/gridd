// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package proto

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize(packet []byte) error
}
