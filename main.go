package main

import "fmt"
import "sort"

//import "huffman/huffheap"
import "container/heap"

//import "huffmantree/huffmantree"

type TreeNode struct {
	left      *TreeNode
	symbol    string
	frequency int
	right     *TreeNode
}

func NewTreeNode(symbol string, freq int) *TreeNode {
	return &TreeNode{symbol: symbol, frequency: freq, left: nil, right: nil}
}

func (node *TreeNode) Left() *TreeNode {
	return node.left
}

func (node *TreeNode) SetLeft(left *TreeNode) {
	node.left = left
}

func (node *TreeNode) Right() *TreeNode {
	return node.right
}

func (node *TreeNode) SetRight(right *TreeNode) {
	node.right = right
}

func (node *TreeNode) Symbol() string {
	return node.symbol
}

func (node *TreeNode) Frequency() int {
	return node.frequency
}

type FreqHeapItem struct {
	value    TreeNode
	priority int
	index    int
}

func NewFreqHeapItem(value TreeNode, priority int) *FreqHeapItem {
	return &FreqHeapItem{value: value, priority: priority}
}

func NewFreqHeapItemWithIndex(value TreeNode, priority int, index int) *FreqHeapItem {
	return &FreqHeapItem{value: value, priority: priority, index: index}
}

func (item *FreqHeapItem) Value() TreeNode {
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
func (mh *MinHeap) Update(item *FreqHeapItem, value TreeNode, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(mh, item.index)
}

//type Encoding map[string]string

func GetHuffmanEncodingMap(hf *TreeNode) map[string]string {
	encoding := make(map[string]string)

	TraverseHuffmanTree(hf, "", encoding)

	return encoding
}

func TraverseHuffmanTree(n *TreeNode, prefix string, encoding map[string]string) {
	if n.Symbol() != "" {
		encoding[n.Symbol()] = prefix
	} else {
		TraverseHuffmanTree(n.Left(), prefix+"0", encoding)
		TraverseHuffmanTree(n.Right(), prefix+"1", encoding)
	}
}

func ComputeHuffmanEncoding(freqs map[string]int) map[string]string {
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(MinHeap, len(freqs))
	i := 0
	for value, priority := range freqs {
		leafNode := NewTreeNode(value, priority)
		pq[i] = NewFreqHeapItemWithIndex(*leafNode, priority, i)
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		item1 := heap.Pop(&pq).(*FreqHeapItem)
		item2 := heap.Pop(&pq).(*FreqHeapItem)

		node1 := item1.Value()
		node2 := item2.Value()
		sum := node1.Frequency() + node2.Frequency()
		newParent := NewTreeNode("", sum)
		newParent.SetLeft(&node1)
		newParent.SetRight(&node2)

		newItem := NewFreqHeapItem(*newParent, sum)
		heap.Push(&pq, newItem)
	}

	item := heap.Pop(&pq).(*FreqHeapItem)
	root := item.Value()

	encoding := GetHuffmanEncodingMap(&root)

	return encoding
}

func ComputeFrequencies(input string) map[string]int {
	freqs := make(map[string]int)

	for _, ch := range input {
		key := string(ch)

		_, ok := freqs[key]
		if ok {
			freqs[key] += 1
		} else {
			freqs[key] = 1
		}
	}

	return freqs
}

func main() {
	input := "I vividly remember one of the first software systems I was responsible for developing and operating. I was fortunate to be given the responsibility of everything from design to support. Over the span of several years, there was one particular sub system that had grown popular. Within the sub system complexity arose from a steady stream of requests to add capabilities. With every change, I meticulously verified existing functionality was not lost. The manual verification was time consuming and became exponentially more complex with each added feature. Like many, I had heard of automated testing but never had the time to learn the techniques. At some point, in my spare time, I happened to pick up a few books and began studying the techniques. Eager to apply this to something real, I realized that automated testing was a perfect fit for the pain I had with the sub system that was growing exponentially difficult to change. One afternoon, I found myself on a road trip with several hours to spare. I had no Internet connection, but I had all the tools I needed to start writing tests. Several days later I had automated the majority of the minutiae I spent hours manually checking. The time savings were insane, several hours of testing turned into several seconds. Truly though, the time savings were secondary. By taking all of those scenarios out of my mind, I experienced an indelible transformation, a restored sense of confidence in what I was creating! Much like the confidence of developing a new system that has yet to mature any complexity. I used to worry excessively about releasing updates. I checked, double checked and triple checked my work. But no matter how much I checked, I still worried. Being responsible for releasing and supporting the application, I innately resisted releases. With a suite of automated checks to lean on, the fear vanished."
	freqs := ComputeFrequencies(input)

	fmt.Println("FREQUENCIES")
	fmt.Println("--------------------------------------------")
	fkeys := []string{}
	for k, _ := range freqs {
		fkeys = append(fkeys, k)
	}
	sort.Strings(fkeys)
	for _, k := range fkeys {
		fmt.Printf("'%s' = %d\n", k, freqs[k])
	}

	encoding := ComputeHuffmanEncoding(freqs)

	fmt.Println("HUFFMAN ENCODING")
	fmt.Println("-------------------------------------------")

	keys := []string{}
	for k, _ := range encoding {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("'%s' = %s\n", k, encoding[k])
	}
}
