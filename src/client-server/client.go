package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/dipankardas011/Efficient-client-server/client"
	"github.com/dipankardas011/Efficient-client-server/payload"
	"github.com/dipankardas011/Efficient-client-server/server"
)

func main() {
	fmt.Println("Hello from [[client]]")
	fmt.Println(`
* GET
* <message>`)
	c, err := net.Dial("tcp", payload.SERVER_HOST+":"+payload.SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the message to be entered..")
		var msg string
		msg, _ = reader.ReadString('\n')
		msg = strings.Replace(msg, "\n", "", -1)
		if strings.Compare(msg, "END") == 0 {
			fmt.Println("Client is terminated!")
			return
		}
		// client.Main_client get the encoded payload to be sent
		encodedMessage, err := client.MainEncoder(msg)
		if err != nil {
			panic(err)
		}

		_, err = c.Write(encodedMessage)
		if err != nil {
			panic(err)
		}

		buffer := make([]byte, 2048)
		mLen, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		fmt.Println("Message from [[server]]: ", server.MainDecoder(buffer[:mLen]))

		// fmt.Println(server.MainDecoder(buffer[:mLen]))
	}
}
