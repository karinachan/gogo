package main

import (
        "fmt"
        "net"
)

type GoClient struct {
        input chan string
        output chan string
        read_write *bufio.ReadWriter
}


// see type Conn for connection info
// essentially can read and write data from connection
func NewGoClient(conn net.Conn) *GoClient {
        rw := bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn))
        goClient := &GoClient{
                input: make(chan string),
                output: make(chan string),
                read_write: rw,
        }
        goClient.Listen()
        return goClient;
}

func (client *GoClient) Read() {
        for { // read line by line
                in, err := client.read_write.Reader.ReadString('\n') 
                if err != nil {
                        fmt.Println("ERROR")
                }
                client.input <-in
        }
}

func (client *GoClient) Write() {
        for current := range client.output {
                client.read_write.Writer.WriteString(current)
                // have to flush, like java
                client.Writer.Flush()
        }
}

func (client *GoClient) Listen() {
        go client.Read()
        go client.Write()
}

func main() {
        fmt.Println("Help")
        hostName := "localhost" // change this
        portNum := "8000"

        conn, err := net.Dial("tcp", hostName+":"+portNum)
        fmt.Println("hehe")
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Println("hi")
        fmt.Println("Connection established between %s and localhost.\n", hostName)
        fmt.Println("Remote Address : %s \n", conn.RemoteAddr().String())
        fmt.Println("Local Address : %s \n", conn.LocalAddr().String())

        goClient := NewGoClient(conn)


}