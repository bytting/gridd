// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package main

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize(packet []byte) error
}
