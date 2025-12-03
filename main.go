package main

import (
	"container/heap"
	"fmt"
)

// ------------------------------------------------------------
// DAY 1: SORTING ALGORITHMS
// ------------------------------------------------------------

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// ------------------------------------------------------------
// DAY 1: STRING OPERATIONS
// ------------------------------------------------------------

func ReverseString(s string) string {
	r := []rune(s)
	l, rIndex := 0, len(r)-1
	for l < rIndex {
		r[l], r[rIndex] = r[rIndex], r[l]
		l++
		rIndex--
	}
	return string(r)
}

func IsPalindrome(s string) bool {
	r := []rune(s)
	l, rIndex := 0, len(r)-1
	for l < rIndex {
		if r[l] != r[rIndex] {
			return false
		}
		l++
		rIndex--
	}
	return true
}

func CountVowelsConsonants(s string) (int, int) {
	vowels := "aeiouAEIOU"
	v, c := 0, 0
	for _, ch := range s {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			if containsRune(vowels, ch) {
				v++
			} else {
				c++
			}
		}
	}
	return v, c
}

func containsRune(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------
// DAY 2: HASH & MAPS
// ------------------------------------------------------------

func WordFrequency(text string) map[string]int {
	freq := make(map[string]int)
	word := ""
	for _, ch := range text {
		if ch == ' ' {
			freq[word]++
			word = ""
		} else {
			word += string(ch)
		}
	}
	if word != "" {
		freq[word]++
	}
	return freq
}

type HashSet struct {
	set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{set: make(map[int]bool)}
}
func (h *HashSet) Add(v int)           { h.set[v] = true }
func (h *HashSet) Contains(v int) bool { return h.set[v] }
func (h *HashSet) Remove(v int)        { delete(h.set, v) }

// ------------------------------------------------------------
// DAY 3: BINARY SEARCH TREE (BST)
// ------------------------------------------------------------

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(value int) {
	bst.root = insertHelper(bst.root, value)
}

func insertHelper(n *Node, value int) *Node {
	if n == nil {
		return &Node{value: value}
	}
	if value < n.value {
		n.left = insertHelper(n.left, value)
	} else if value > n.value {
		n.right = insertHelper(n.right, value)
	}
	return n
}

func (bst *BST) Search(value int) bool {
	return searchHelper(bst.root, value)
}

func searchHelper(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	}
	if value < n.value {
		return searchHelper(n.left, value)
	}
	return searchHelper(n.right, value)
}

func findMin(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (bst *BST) Delete(value int) {
	bst.root = deleteHelper(bst.root, value)
}

func deleteHelper(n *Node, value int) *Node {
	if n == nil {
		return nil
	}
	if value < n.value {
		n.left = deleteHelper(n.left, value)
	} else if value > n.value {
		n.right = deleteHelper(n.right, value)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		minNode := findMin(n.right)
		n.value = minNode.value
		n.right = deleteHelper(n.right, minNode.value)
	}
	return n
}

func (bst *BST) Inorder() {
	inorderHelper(bst.root)
	fmt.Println()
}

func inorderHelper(n *Node) {
	if n != nil {
		inorderHelper(n.left)
		fmt.Print(n.value, " ")
		inorderHelper(n.right)
	}
}

// ------------------------------------------------------------
// DAY 4: MIN HEAP (Custom)
// ------------------------------------------------------------

type MinHeap struct {
	data []int
}

func (h *MinHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MinHeap) left(i int) int   { return 2*i + 1 }
func (h *MinHeap) right(i int) int  { return 2*i + 2 }

func (h *MinHeap) Insert(v int) {
	h.data = append(h.data, v)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.data[i] < h.data[h.parent(i)] {
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i)
	}
}

func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap empty")
	}
	min := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyDown(i int) {
	smallest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.data) && h.data[l] < h.data[smallest] {
		smallest = l
	}
	if r < len(h.data) && h.data[r] < h.data[smallest] {
		smallest = r
	}
	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapifyDown(smallest)
	}
}

// ------------------------------------------------------------
// DAY 4: MAX HEAP (Custom)
// ------------------------------------------------------------

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MaxHeap) left(i int) int   { return 2*i + 1 }
func (h *MaxHeap) right(i int) int  { return 2*i + 2 }

func (h *MaxHeap) Insert(v int) {
	h.data = append(h.data, v)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 && h.data[i] > h.data[h.parent(i)] {
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i)
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap empty")
	}
	max := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return max
}

func (h *MaxHeap) heapifyDown(i int) {
	largest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.data) && h.data[l] > h.data[largest] {
		largest = l
	}
	if r < len(h.data) && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.heapifyDown(largest)
	}
}

// ------------------------------------------------------------
// DAY 4: PRIORITY QUEUE USING container/heap
// ------------------------------------------------------------

type Item struct {
	value    string
	priority int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// ------------------------------------------------------------
// MAIN - RUN DSA TOOLKIT DEMO
// ------------------------------------------------------------

func main() {
	fmt.Println("=== Sorting ===")
	a := []int{5, 2, 9, 1}
	BubbleSort(a)
	fmt.Println("Bubble:", a)

	b := []int{7, 4, 1, 9}
	InsertionSort(b)
	fmt.Println("Insertion:", b)

	c := []int{6, 2, 8, 3}
	SelectionSort(c)
	fmt.Println("Selection:", c)

	fmt.Println("\n=== Strings ===")
	fmt.Println("Reverse:", ReverseString("malayalam"))
	fmt.Println("Palindrome:", IsPalindrome("malayalam"))
	v, con := CountVowelsConsonants("Hello World")
	fmt.Println("Vowels:", v, "Consonants:", con)

	fmt.Println("\n=== Word Frequency ===")
	fmt.Println(WordFrequency("apple banana apple mango mango grape"))

	fmt.Println("\n=== HashSet ===")
	set := NewHashSet()
	set.Add(10)
	set.Add(20)
	fmt.Println("Contains 10:", set.Contains(10))

	fmt.Println("\n=== BST ===")
	bst := &BST{}
	bst.Insert(8)
	bst.Insert(3)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(6)
	fmt.Print("Inorder: ")
	bst.Inorder()
	fmt.Println("Search 6:", bst.Search(6))
	bst.Delete(3)
	fmt.Print("After Delete: ")
	bst.Inorder()

	fmt.Println("\n=== Min Heap ===")
	minH := &MinHeap{}
	minH.Insert(10)
	minH.Insert(3)
	minH.Insert(1)
	minH.Insert(7)
	fmt.Println("Heap:", minH.data)
	fmt.Println("Pop:", minH.Pop())

	fmt.Println("\n=== Max Heap ===")
	maxH := &MaxHeap{}
	maxH.Insert(10)
	maxH.Insert(3)
	maxH.Insert(15)
	maxH.Insert(7)
	fmt.Println("Heap:", maxH.data)
	fmt.Println("Pop:", maxH.Pop())

	fmt.Println("\n=== Priority Queue ===")
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: "task1", priority: 2})
	heap.Push(pq, &Item{value: "task2", priority: 5})
	heap.Push(pq, &Item{value: "task3", priority: 1})

	fmt.Println("Pop Highest Priority:", heap.Pop(pq).(*Item))
}
