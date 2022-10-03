package client

import (
	"bufio"
	"fmt"
	"github.com/dipankardas011/Efficient-client-server/heap"
	"github.com/dipankardas011/Efficient-client-server/payload"
	"os"
	"strings"
)

var (
	HeapSize int
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "9988"
)

func GetFreq(message string) map[byte]uint64 {
	freq := make(map[byte]uint64)

	for _, character := range message {
		freq[byte(character)] = freq[byte(character)] + 1
	}
	return freq
}

func GenerateHeap(message string, hashMap map[byte]uint64) heap.HeapArrDS {
	heapArr := make(heap.HeapArrDS, len(hashMap)+1)
	index := 1
	for key, value := range hashMap {
		heapArr[index] = new(heap.HeapDS)
		heapArr[index].Character = key
		heapArr[index].Left = nil
		heapArr[index].Right = nil
		heapArr[index].Frequency = int64(value)
		index++
	}
	HeapSize = index
	return heapArr
}

func getHuffmanCodes(root *heap.HeapDS, encode string, table *map[byte]string) {
	if root.Left == nil && root.Right == nil {
		(*table)[root.Character] = encode
		return
	}
	getHuffmanCodes(root.Left, encode+"0", table)
	getHuffmanCodes(root.Right, encode+"1", table)
}

func Main_client() payload.Payload {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the message to be entered..")
	var msg string
	msg, _ = reader.ReadString('\n')
	msg = strings.Replace(msg, "\n", "", -1)
	hashMap := GetFreq(msg)

	var heapArr heap.HeapArrDS
	heapArr = GenerateHeap(msg, hashMap)

	heapArrPtr := &heapArr
	heapArrPtr.BuildHeap(HeapSize)

	for HeapSize > 2 {
		var x, y, z *heap.HeapDS
		x = heapArrPtr.PopHeap(&HeapSize)
		y = heapArrPtr.PopHeap(&HeapSize)

		z = new(heap.HeapDS)
		z.Frequency = x.Frequency + y.Frequency
		z.Left = x
		z.Right = y
		heapArrPtr.PushHeap(z, &HeapSize)
	}
	var root *heap.HeapDS
	root = heapArrPtr.PopHeap(&HeapSize)
	var tableHeap map[byte]string
	tableHeap = make(map[byte]string)
	getHuffmanCodes(root, "", &tableHeap)

	var encodedMsg string
	for _, char := range msg {
		encodedMsg += tableHeap[byte(char)]
	}

	///// Add code to send the table and the encoded message
	var ret payload.Payload
	ret = payload.PayloadDS{}
	ret = ret.AddInfo(encodedMsg, tableHeap)
	return ret
}