package p2p

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// CNodeManager structure
type CNodeManager struct {
	connections map[*CNode]bool
	broadcast   chan []byte
	register    chan *CNode
	unregister  chan *CNode
}

// CNode Structure
type CNode struct {
	socket net.Conn
	data   chan []byte
}

func (manager *CNodeManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.connections[connection] = true
			fmt.Printf("[+] Established new connection <%s>\n", connection.socket.RemoteAddr().String())
		case connection := <-manager.unregister:
			if _, ok := manager.connections[connection]; ok {
				close(connection.data)
				delete(manager.connections, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.connections {
				select {
				case connection.data <- message:
				default:
						close(connection.data)
						delete(manager.connections, connection)
					}
				}
		}
	}
}

func (manager *CNodeManager) receive(client *CNode) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("[RECEIVED]: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (client *CNode) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("[RECEIVED]: " + string(message))
		}
	}
}

func (manager *CNodeManager) send(client *CNode) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func startServerMode() {
	fmt.Println("Hosting the server...")
	listener, error := net.Listen("tcp", ":12345")
	if error != nil {
		fmt.Println(error)
	}
	manager := CNodeManager{
		connections:    make(map[*CNode]bool),
		broadcast:  		make(chan []byte),
		register:   		make(chan *CNode),
		unregister: 		make(chan *CNode),
	}
	go manager.start()
	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &CNode{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}

func startClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", "localhost:12345")
	if error != nil {
		fmt.Println(error)
	}
	client := &CNode{socket: connection}
	go client.receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

// Start p2p network
func Start(mode string) {
	if mode == "server" {
		startServerMode()
	} else {
		startClientMode()
	}
}
