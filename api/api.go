// LICENSE: GNU General Public License version 2
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (go.libremail AT gmail DOT com)

package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

/*

Json API request examples:

{"Request":"connect","Args":["127.0.0.1:30000"]}
{"Request":"quit","Args":null}
{"Request":"list","Args":["addresses","peers"]}

Json API reply examples:

{"Reply":"addresses","Items":[{"address":"LM:1234567890"},{"address":"LM:0987654321"}]}

*/

type RequestType struct {
	Request string
	Args    []string
}

func (r *RequestType) AddArgument(arg string) *RequestType {

	r.Args = append(r.Args, arg)
	return r
}

func (r *RequestType) String() string {

	b, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}

type ReplyType struct {
	Reply string
	Items []map[string]string
}

func (r *ReplyType) String() string {

	b, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}

type JsonService struct {
	RequestChan chan *RequestType
	ReplyChan   chan *ReplyType
}

func (js *JsonService) Run(serviceGroup *sync.WaitGroup) {

	log.Println("Starting Json service")

	serviceGroup.Add(1)
	defer serviceGroup.Done()

	cs := &consoleService{make(chan string)}
	go cs.Run()

	for {
		select {

		case cmd := <-cs.CommandChan:

			req := new(RequestType)
			if err := json.Unmarshal([]byte(cmd), &req); err != nil {
				fmt.Println(err)
				continue
			}

			js.RequestChan <- req

			// If we have a quit command, exit this service
			if req.Request == "quit" {

				log.Println("Quitting json service")
				return
			}

		case rep := <-js.ReplyChan:

			fmt.Println(rep)
		}
	}
}

type consoleService struct {
	CommandChan chan string
}

func (cs *consoleService) Run() {

	log.Println("Starting console service")

	// Read commands from stdin
	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {

			fmt.Println("ERROR:", err)
			break
		}

		cmd := strings.Trim(line, "\n\r\t ")

		if len(cmd) > 0 {
			cs.CommandChan <- cmd
		}

		// If we have a quit command, exit this service
		if strings.HasPrefix(cmd, "{\"Identifier\":\"quit\"") {

			log.Println("Quitting console service")
			return
		}
	}
}
