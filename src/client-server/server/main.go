package server

import (
	"encoding/json"
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/payload"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "9988"
)

func DecodeMessage(encoded payload.PayloadDS) string {
	var reverseHashMap map[string]byte
	reverseHashMap = make(map[string]byte, len(encoded.GetTable()))
	for key, value := range encoded.GetTable() {
		reverseHashMap[value] = key
	}
	i := 0
	j := 1
	decodedMsg := ""
	for i < len(encoded.GetEncoded()) && j <= len(encoded.GetEncoded()) {
		if value, found := reverseHashMap[encoded.GetEncoded()[i:j]]; found {
			decodedMsg += string(value)
			i = j
		}
		j++
	}
	return decodedMsg
}

func Main_server(jsonRecv []byte) {
	// wait for the data
	fmt.Println("Data processing for ", string(jsonRecv))
	recvData := payload.PayloadDS{}
	err := json.Unmarshal(jsonRecv, &recvData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message from [[CLIENT]] -> ", DecodeMessage(recvData))
}
