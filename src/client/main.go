package client

import (
	"encoding/json"
	"github.com/dipankardas011/Efficient-client-server/heap"
	"github.com/dipankardas011/Efficient-client-server/payload"
)

var (
	HeapSize int
)

// MainEncoder FIXME: Minor buffer full incase of creating the encoded message if message is to big it does not fit into the []byte for transmittion
// MainEncoder handler of message encoding and decoding
func MainEncoder(msg string) ([]byte, error) {
	hashMap := heap.GetFreq(msg)

	var heapArr heap.HeapArrDS
	heapArr = heap.GenerateHeap(hashMap, &HeapSize)

	var root *heap.HeapDS
	root = heap.GenerateHuffmanTree(&heapArr, &HeapSize)

	var tableHeap map[byte]string
	tableHeap = make(map[byte]string)
	heap.GetHuffmanCodes(root, "", &tableHeap)

	var encodedMsg string
	for _, char := range msg {
		encodedMsg += tableHeap[byte(char)]
	}

	var ret payload.Payload
	ret = payload.PayloadDS{}
	ret = ret.AddInfo(encodedMsg, tableHeap)
	byteArray, err := json.Marshal(ret)
	return byteArray, err
}
