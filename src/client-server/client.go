package main

import (
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/client"
	"net"
)

func main() {
	fmt.Println("Hello from [[client]]")
	choice := 1
	c, err := net.Dial("tcp", client.SERVER_HOST+":"+client.SERVER_PORT)
	if err != nil {
		panic(err)
	}
	for choice != 0 {

		// client.Main_client get the encoded payload to be sent
		encodedMessage, err := client.Main_client()
		if err != nil {
			panic(err)
		}

		_, err = c.Write(encodedMessage)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n\nWhether to continue [1/0]")
		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			panic("Choice Err!")
		}
	}
	_, err = c.Write([]byte("END"))
	if err != nil {
		panic(err)
	}
	c.Close()
}
