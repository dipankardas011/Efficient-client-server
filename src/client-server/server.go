package main

import (
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/server"
	"net"
	"strings"
)

func main() {
	fmt.Println("Hello from [[server]]")

	l, err := net.Listen("tcp", server.SERVER_HOST+":"+server.SERVER_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Listening on ", server.SERVER_HOST+":"+server.SERVER_PORT)
	defer l.Close()
	defer fmt.Println("Server terminated!!")

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for true {
		buffer := make([]byte, 2048)
		mLen, err := c.Read(buffer)

		if strings.Compare("END", strings.TrimSpace(string(buffer[:mLen]))) == 0 {
			break
		}

		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		server.Main_server(buffer[:mLen])
		buffer = nil
	}
}
