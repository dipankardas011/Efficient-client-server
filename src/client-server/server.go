package main

import (
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/payload"
	"github.com/dipankardas011/Efficient-client-server/server"
	"net"
	"strings"
)

func main() {
	fmt.Println("Hello from [[server]]")

	l, err := net.Listen("tcp", payload.SERVER_HOST+":"+payload.SERVER_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Listening on ", payload.SERVER_HOST+":"+payload.SERVER_PORT)
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// server is always open
	for true {
		buffer := make([]byte, 2048)
		mLen, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Message from buffer: ", strings.TrimSpace(string(buffer[:mLen])))

		choice := server.MainDecoder(buffer[:mLen])

		switch choice {
		case "GET":
			fmt.Println("\tGet the index.html!")
		default:
			fmt.Println("\tDefault is called!")
		}

		buffer = nil
	}
}
