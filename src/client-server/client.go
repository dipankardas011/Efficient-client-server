package main

import (
	"encoding/json"
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/client"
	"github.com/dipankardas011/Efficient-client-server/payload"
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

		var clientPayload payload.Payload
		clientPayload = client.Main_client()

		//fmt.Println(clientPayload.GetEncoded())
		//fmt.Println(clientPayload.GetTable())
		byteArray, err := json.Marshal(clientPayload)
		_, err = c.Write(byteArray)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(byteArray))
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
