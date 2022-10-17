package server

import (
	"encoding/json"

	"github.com/dipankardas011/Efficient-client-server/payload"
)

func MainDecoder(jsonRecv []byte) string {
	// wait for the data
	// fmt.Println("Data processing for ", string(jsonRecv))
	recvData := payload.PayloadDS{}
	err := json.Unmarshal(jsonRecv, &recvData)
	if err != nil {
		panic(err)
	}
	ret := payload.DecodeMessage(recvData)
	return ret
}
