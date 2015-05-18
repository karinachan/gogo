/* GoClient.go

This file includes a main that creates a connection to a remote server

*/

package main

import (
	"fmt"
        "bufio"
	"net"
        "os"
)


func main() {
        hostName := "localhost" // change this
        portNum := "8000"

        conn, err := net.Dial("tcp", hostName+":"+portNum)

        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Println("Connection established between Client and localhost.")

        // There is a misunderstanding about how to channel data to the server. 
        // Assumedly, 'conn' is the channel produced between client and server.

        read := bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(os.Stdout))
        write := bufio.NewReadWriter(bufio.NewReader(os.Stdin),bufio.NewWriter(conn))

        for {
                go run(read)
                go run(write)
        }       
}

func run(rw *bufio.ReadWriter) {
        for {
                line, _ := rw.Reader.ReadString('\n')
                rw.Writer.WriteString(line)
                fmt.Println(line) // even this isn't working. So I don't really know what's up.
        }
}
