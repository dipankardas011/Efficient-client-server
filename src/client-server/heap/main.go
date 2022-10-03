package heap

import "fmt"

type HeapDS struct {
	Character byte
	Frequency int64
	Left      *HeapDS
	Right     *HeapDS
}

const HEAP_ARR_SIZE = 100

func (h *HeapArrDS) Swap(l, r int) {
	var temp *HeapDS
	temp = (*h)[l]
	(*h)[l] = (*h)[r]
	(*h)[r] = temp
}

func (h *HeapArrDS) Heapify(index, size int) {
	l := 2 * index
	r := 2*index + 1
	smallest := index
	if l < size && (*h)[l].Frequency < (*h)[smallest].Frequency {
		smallest = l
	}
	if r < size && (*h)[r].Frequency < (*h)[smallest].Frequency {
		smallest = r
	}
	if smallest == index {
		return
	}
	h.Swap(smallest, index)
	h.Heapify(smallest, size)
}

func (h *HeapArrDS) BuildHeap(HeapSize int) {
	for i := HeapSize / 2; i >= 1; i-- {
		h.Heapify(i, HeapSize)
	}
}

func (h *HeapArrDS) PopHeap(HeapSize *int) *HeapDS {
	var abcd *HeapDS
	abcd = (*h)[1]
	(*h)[1] = (*h)[*HeapSize-1]
	*HeapSize = *HeapSize - 1
	h.Heapify(1, *HeapSize)

	return abcd
}

func (h *HeapArrDS) PushHeap(item *HeapDS, HeapSize *int) {
	*HeapSize = *HeapSize + 1
	i := *HeapSize - 1
	for i > 1 && item.Frequency < (*h)[i/2].Frequency {
		(*h)[i] = (*h)[i/2]
		i /= 2
	}
	(*h)[i] = item
}

func (h *HeapArrDS) PeekHeap(size int) *HeapDS {
	if size > 0 {
		return (*h)[1]
	}
	return nil
}

type HeapArrDS []*HeapDS

func (h HeapArrDS) Display(HeapSize int) {
	for i := 1; i < HeapSize; i++ {
		fmt.Println(*(h[i]))
	}
}

type Heap interface {
	BuildHeap(int)
	Heapify(int, int)
	Display(int)
	Swap(int, int)
	PopHeap(*int) *HeapDS
	PushHeap(*HeapDS, *int)
	PeekHeap(int) *HeapDS
}
