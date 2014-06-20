// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag.robole AT gmail DOT com)

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"gridd/api"
	"gridd/grid"
	"gridd/proto"
	"gridd/store"
)

func main() {

	// Parse commandline flags
	port := flag.Uint("port", 30000, "The listening port")
	logfile := flag.String("logfile", "log.txt", "The log file")
	flag.Parse()

	if flag.NFlag() < 2 {
		fmt.Println("Usage:", os.Args[0], "-port=... -logfile=...")
		return
	}

	// Setup logfile
	logfd, err := os.Create(*logfile)
	if err != nil {
		panic(err)
	}
	defer logfd.Close()
	log.SetOutput(logfd)

	// Temporary map to hold peer connections
	peers := make(map[string]net.Conn)

	// WaitGroup to synchronize service shutdown
	serviceGroup := &sync.WaitGroup{}

	// Open databases
	db := store.NewStore("./private.db", "./public.db")
	if err = db.Open(); err != nil {
		panic(err)
	}
	defer db.Close()

	addr1, err := proto.NewAddress(1, 0)
	if err != nil {
		panic(err)
	}
	db.SaveAddress(addr1)

	// Start Json API service
	apiService := &api.JsonService{make(chan *api.RequestType), make(chan *api.ReplyType)}
	go apiService.Run(serviceGroup)

	// Start listen service
	listenService := &grid.ListenService{uint16(*port), make(chan net.Conn), make(chan struct{})}
	go listenService.Run(serviceGroup)

	// Start connect service
	connectService := &grid.ConnectService{make(chan string), make(chan net.Conn)}
	go connectService.Run(serviceGroup)

	// Start handshake service
	handshakeService := &grid.HandshakeService{make(chan net.Conn)}
	go handshakeService.Run(serviceGroup)

	// Start initiate handshake service
	initiateHandshakeService := &grid.InitiateHandshakeService{make(chan net.Conn)}
	go initiateHandshakeService.Run(serviceGroup)

L1:
	for { // Event loop

		select {
		case req := <-apiService.RequestChan:

			log.Println("Received command", req.Request)

			switch req.Request {

			case "quit":

				break L1

			case "connect":

				if len(req.Args) > 0 {
					connectService.AddressChan <- req.Args[0]
				}

			case "list":

				switch req.Args[0] {

				case "peers":

					rep := new(api.ReplyType)
					rep.Reply = "peers"

					for i := 0; i < len(peers); i++ {
						rep.Items[i] = make(map[string]string)
					}

					i := 0
					for k, _ := range peers {
						rep.Items[i]["ip"] = k
						i++
					}

					apiService.ReplyChan <- rep

				case "addresses":

					// FIXME: faking a few addresses
					rep := new(api.ReplyType)
					rep.Reply = "addresses"

					rep.Items = append(rep.Items, make(map[string]string))
					rep.Items = append(rep.Items, make(map[string]string))

					rep.Items[0]["address"] = "LM:1234567890"
					rep.Items[1]["address"] = "LM:0987654321"

					apiService.ReplyChan <- rep

				default:
				}

			default:
			}

		case connection := <-listenService.ConnectionChan:

			handshakeService.ConnectionChan <- connection

		case connection := <-handshakeService.ConnectionChan:

			peers[connection.RemoteAddr().String()] = connection
			//db.RegisterPeer(connection)

		case connection := <-connectService.ConnectionChan:

			initiateHandshakeService.ConnectionChan <- connection

		case connection := <-initiateHandshakeService.ConnectionChan:

			peers[connection.RemoteAddr().String()] = connection
			//db.RegisterPeer(connection)
		}
	}

	log.Println("Stopping services")

	listenService.Close()
	connectService.Close()
	handshakeService.Close()
	initiateHandshakeService.Close()

	serviceGroup.Wait()

	for k, v := range peers {

		log.Println("Closing peer", k)
		v.Close()
	}
}
