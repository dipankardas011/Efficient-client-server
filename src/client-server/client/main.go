package client

import (
	"bufio"
	"encoding/json"
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

// GetFreq return the hashMap of frequency of each character
func GetFreq(message string) map[byte]uint64 {
	freq := make(map[byte]uint64)

	for _, character := range message {
		freq[byte(character)] = freq[byte(character)] + 1
	}
	return freq
}

// GenerateHeap returns HeapArray (MinHeap) Datastructure from hashMap from GetFreq
func GenerateHeap(hashMap map[byte]uint64) heap.HeapArrDS {
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

// getHuffmanCodes Generate the Huffman codes by Tree Traversal
func getHuffmanCodes(root *heap.HeapDS, encode string, table *map[byte]string) {

	if root != nil {
		if root.Left == nil && root.Right == nil {
			(*table)[root.Character] = encode
			return
		}
		getHuffmanCodes(root.Left, encode+"0", table)
		getHuffmanCodes(root.Right, encode+"1", table)
	}
}

// generateHuffmanTree create the huffmanTree using the MinHeapArray
func generateHuffmanTree(heapArrPtr *heap.HeapArrDS) *heap.HeapDS {
	heapArrPtr.BuildHeap(HeapSize)
	if HeapSize == 2 {
		// only one element is there
		var x, z *heap.HeapDS
		x = heapArrPtr.PopHeap(&HeapSize)
		z = new(heap.HeapDS)
		z.Frequency = x.Frequency
		z.Left = x
		z.Right = nil
		heapArrPtr.PushHeap(z, &HeapSize)
	}

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
	return heapArrPtr.PopHeap(&HeapSize)
}

// Main_client handler of message encoding and decoding
func Main_client() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the message to be entered..")
	var msg string
	msg, _ = reader.ReadString('\n')
	msg = strings.Replace(msg, "\n", "", -1)
	hashMap := GetFreq(msg)

	var heapArr heap.HeapArrDS
	heapArr = GenerateHeap(hashMap)

	var root *heap.HeapDS
	root = generateHuffmanTree(&heapArr)

	var tableHeap map[byte]string
	tableHeap = make(map[byte]string)
	getHuffmanCodes(root, "", &tableHeap)

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
