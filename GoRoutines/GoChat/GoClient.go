/* GoClient.go

This file includes a main that creates a connection to a remote server

*/

package main

import (
	"fmt"
	"net"
)

func main() {
        hostName := "localhost" // change this
        portNum := "8000"

        conn, err := net.Dial("tcp", hostName+":"+portNum)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Println("Connection established between %s and localhost.\n", hostName)
        fmt.Println("Remote Address : %s \n", conn.RemoteAddr().String())
        fmt.Println("Local Address : %s \n", conn.LocalAddr().String())

}