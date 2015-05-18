/* GoCLient.go
 Author: Emery Otopalik

 Sources: 
 https://gist.github.com/drewolson/3950226
 https://www.socketloop.com/tutorials/golang-simple-client-server-example
 Basically just used sources to gather info on what packages to use/how to use/syntax.

 This program creates a chat server that manages multiple clients. Input from clients
 is sent to the server, which in following sends the data to all clients.*/

package main

import (
	"fmt" // for printing 
	"bufio" // for reading/writing
	"net" // for connecting to host, managing client/server connection(s)
)

// Data structure for a server.
type GoServer struct {
	goClients [] *GoClient // array to manage clients
	add chan net.Conn // channel for managing connections between server and clients
	input chan string // input channel
	output chan string // output channel
}

// Creates a new chat server. The server is set to continuously listen to clients.
func NewGoServer() *GoServer {
	goServer := &GoServer {
		goClients: make([]*GoClient, 0),
		add: make(chan net.Conn),
		input: make(chan string),
		output: make(chan string),
	}
	goServer.Listen()
	return goServer
}

// Adds a new client to the server. Server is added to client array and channels are 
// pointed appropriately.
func (goServer *GoServer) Add(conn net.Conn) {
	goClient := NewGoClient(conn)
	goServer.goClients = append(goServer.goClients, goClient)
	// client's input sent to server to be inputted to other clients
	go func() { for  // client's input will be passed to server's input
		{ goServer.input<- <-goClient.input } 
	}()
	//goClient.Run() 
}

// The data stored in the server output channel is sent to each client.
func (goServer *GoServer) Output(out string) {
	for _, client := range goServer.goClients {	
		client.output <- out
	}
}

// The server listens for both new input and new clients. If new input is found, it is sent
// to the server output. If a new client is found, it is added to the server.
func (goServer *GoServer) Listen() {
	go func() { 
		for { 
			select { // select statement to choose 'thread'
				case out := <-goServer.input: goServer.Output(out)
				case conn := <-goServer.add: goServer.Add(conn)
			}
		}
	}()
}

// Data structure for a client.
type GoClient struct {
	input chan string // input channel
	output chan string // output channel
	read_write *bufio.ReadWriter // Read/Write Buffer
}

// Creates a new client, connected to the server. The client begins running.
func NewGoClient(conn net.Conn) *GoClient {
	// a Read/Write buffer
	rw := bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn)) 
	goClient := &GoClient{
		input: make(chan string),
		output: make(chan string),
		read_write: rw,
	}
	goClient.Run()
	return goClient;
}

// When called, the specified client continuously reads incoming input from user and 
// passes it to the input channel.
func (client *GoClient) Read() {
	for { // read line by line
		in, _ := client.read_write.Reader.ReadString('\n') 
		client.input <- in
	}
}

// When called, the specified client continously outputs data from the output channel 
// (from the server/other clients).
func (client *GoClient) Write() {
	for current := range client.output {
		client.read_write.Writer.WriteString(current) 
		client.read_write.Writer.Flush() // flushes any buffering output
	}
}

// Concurrently calls Read() and Write() so that the client is continously doing both.
func (client *GoClient) Run() {
	go client.Read()
	go client.Write()
}


func main() {
	goServer := NewGoServer()

 	ln, err := net.Listen("tcp", ":8000")
 	if err != nil {
 		fmt.Println("ERROR: in net.Listen()")
 	}

 	fmt.Println("Server up and listening on port 8000")

 	for {
 		conn, _ := ln.Accept()
 		if err != nil {
 			fmt.Println("ERROR: in Accept()")
 			continue
 		}

 	goServer.add <- conn

 	}
}