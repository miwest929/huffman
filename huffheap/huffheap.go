package huffheap

import (
	"container/heap"
	"huffman/huffmantree"
)

type FreqHeapItem struct {
	value    string
	priority int
	index    int
}

func NewFreqHeapItem(value string, priority int) *FreqHeapItem {
	return &FreqHeapItem{value: value, priority: priority}
}

func NewFreqHeapItemWithIndex(value string, priority int, index int) *FreqHeapItem {
	return &FreqHeapItem{value: value, priority: priority, index: index}
}

func (item *FreqHeapItem) Value() string {
	return item.value
}

func (item *FreqHeapItem) Priority() int {
	return item.priority
}

func (item *FreqHeapItem) Index() int {
	return item.index
}

type MinHeap []*FreqHeapItem

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return mh[i].priority < mh[j].priority
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MinHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*FreqHeapItem)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*mh = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (mh *MinHeap) Update(item *FreqHeapItem, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(mh, item.index)
}
