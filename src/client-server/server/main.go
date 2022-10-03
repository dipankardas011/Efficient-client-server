package server

import "fmt"

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "9988"
)

func Main_server(jsonRecv []byte) {
	// wait for the data
	fmt.Println("Data processing for ", string(jsonRecv))
}
