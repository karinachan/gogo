package main

import (
	"fmt"
	"log"
	"net"
)

type GoServer struct {
	//clients
	goClients [] *GoClient
	//join channel
	add chan net.Conn
	input chan string
	output chan string
}

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

func (goServer *GoServer) Add(conn net.Conn) {
	goClient := NewGoClient(conn)
	goServer.goClients = append(goServer.goClients, goClient)
	// client's input sent to server to be inputted to other clients
	go func() { for 
		{ goServer.input <- <-goClient.input } 
	}() 
}


func (goServer *GoServer) Output(out string) {
	for _, client := range goServer.goClients {	
		client.output <-out
	}
}

// need to listen to multiple clients at once
// 
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


func handleConnection(c net.Conn) {

 	fmt.Println("Client %v connected.", c.RemoteAddr())



	fmt.Println("Connection from %v closed.", c.RemoteAddr())
 }

func main() {
 	ln, err := net.Listen("tcp", ":8000")
 	if err != nil {
 		log.Fatal(err)
 	}

 	fmt.Println("Server up and listening on port 6000")

 	for {
 		fmt.Println("nono")
 		conn, err := ln.Accept()
 		if err != nil {
 			log.Println(err)
 			continue
 		}
 		go handleConnection(conn)
 	}
}