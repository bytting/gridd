// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package grid

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type ListenService struct {
	Port           uint16
	ConnectionChan chan net.Conn
	CloseChan      chan struct{}
}

func (ls *ListenService) Run(serviceGroup *sync.WaitGroup) {

	log.Println("Starting listen service")

	serviceGroup.Add(1)
	defer serviceGroup.Done()

	laddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", ls.Port))
	if nil != err {

		log.Fatalln("grid.ListenService.Run:", err)
		return
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if nil != err {

		log.Fatalln("grid.ListenService.Run:", err)
		return
	}

	log.Println("Listening on", listener.Addr())

	for {
		// Shutdown gracefully if we have a close signal waiting
		select {
		case <-ls.CloseChan:

			listener.Close()
			return

		default:
		}

		// Accept incoming connections for 3 seconds
		listener.SetDeadline(time.Now().Add(time.Second * 3))

		conn, err := listener.AcceptTCP()
		if err != nil {

			// Accept returns an error on timeout. If we have a timeout, continue
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {

				continue

			} else {

				log.Fatalln("grid.ListenService.Run:", err)
				return
			}
		}

		// Send connection back to client
		ls.ConnectionChan <- conn
	}
}

func (ls *ListenService) Close() {
	close(ls.CloseChan)
}

type ConnectService struct {
	AddressChan    chan string
	ConnectionChan chan net.Conn
}

func (cs *ConnectService) Run(serviceGroup *sync.WaitGroup) {

	log.Println("Starting connect service")

	serviceGroup.Add(1)
	defer serviceGroup.Done()

	for {
		addr, ok := <-cs.AddressChan

		if !ok { return }

		conn, err := net.Dial("tcp", addr)

		if err != nil {

			log.Println("grid.ConnectService.Run:", err)

		} else {

			cs.ConnectionChan <- conn
		}
	}
}

func (cs *ConnectService) Close() {
	close(cs.AddressChan)
}

type HandshakeService struct {
	ConnectionChan chan net.Conn
}

func (hs *HandshakeService) Run(serviceGroup *sync.WaitGroup) {

	log.Println("Starting handshake service")

	serviceGroup.Add(1)
	defer serviceGroup.Done()

	for {
		connection, ok := <-hs.ConnectionChan

		if !ok { return }

		log.Println("Doing handshake with", connection.RemoteAddr())
		// TODO: Handshaking

		hs.ConnectionChan <- connection
	}
}

func (hs *HandshakeService) Close() {
	close(hs.ConnectionChan)
}

type InitiateHandshakeService struct {
	ConnectionChan chan net.Conn
}

func (ihs *InitiateHandshakeService) Run(serviceGroup *sync.WaitGroup) {

	log.Println("Starting initiate handshake service")

	serviceGroup.Add(1)
	defer serviceGroup.Done()

	for {
		connection, ok := <-ihs.ConnectionChan

		if !ok { return }

		log.Println("Initiating handshake with", connection.RemoteAddr())
		// TODO: Handshaking

		ihs.ConnectionChan <- connection
	}
}

func (ihs *InitiateHandshakeService) Close() {
	close(ihs.ConnectionChan)
}
