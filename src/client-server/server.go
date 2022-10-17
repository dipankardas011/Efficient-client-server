package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dipankardas011/Efficient-client-server/client"
	"github.com/dipankardas011/Efficient-client-server/payload"
	"github.com/dipankardas011/Efficient-client-server/server"
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

		choice := server.MainDecoder(buffer[:mLen])
		fmt.Println("Message from [[client]]: ", choice)

		switch choice {
		case "GET":
			fmt.Println("\tGet the index.html!")
			byteFile, err := os.ReadFile("/go/src/server/resources/index.html")
			if err != nil {
				panic(err.Error())
			}

			encodedMessage, err := client.MainEncoder(string(byteFile))
			if err != nil {
				panic(err)
			}

			if _, err = c.Write(encodedMessage); err != nil {
				panic(err)
			}

		default:
			fmt.Println("\tDefault is called!")
			var resp string
			resp = "F** man got the wrong option try with GET"

			encodedMessage, err := client.MainEncoder(resp)
			if err != nil {
				panic(err)
			}

			if _, err = c.Write(encodedMessage); err != nil {
				panic(err)
			}

		}

		buffer = nil
	}
}
