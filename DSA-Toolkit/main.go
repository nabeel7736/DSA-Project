package main

import (
	"container/heap"
	"fmt"

	bst "dsa-toolkit/BST"
	"dsa-toolkit/hashmap"
	dsheap "dsa-toolkit/heap"
	"dsa-toolkit/sorting"
	"dsa-toolkit/strings"
)

func main() {
	// Sortings
	fmt.Println("=== Sorting ===")
	a := []int{5, 2, 9, 1}
	sorting.BubbleSort(a)

	fmt.Println("Bubble:", a)

	b := []int{7, 4, 1, 9}
	sorting.InsertionSort(b)
	fmt.Println("Insertion:", b)

	c := []int{6, 2, 8, 3}
	sorting.SelectionSort(c)
	fmt.Println("Selection:", c)

	fmt.Println("\n=== Strings ===")
	fmt.Println("Reverse:", strings.ReverseString("malayalam"))
	fmt.Println("Palindrome:", strings.IsPalindrome("malayalam"))
	v, con := strings.CountVowelsConsonants("Hello World")
	fmt.Println("Vowels:", v, "Consonants:", con)

	fmt.Println("\n=== Word Frequency ===")
	fmt.Println(hashmap.WordFrequency("apple banana apple mango mango grape"))

	fmt.Println("\n=== HashSet ===")
	set := hashmap.NewHashSet()
	set.Add(10)
	set.Add(20)
	fmt.Println("Contains 10:", set.Contains(10))

	fmt.Println("\n=== BST ===")
	bst := &bst.BST{}
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
	minH := &dsheap.MinHeap{}
	minH.Insert(10)
	minH.Insert(3)
	minH.Insert(1)
	minH.Insert(7)
	fmt.Println("Heap:", minH.Data)
	fmt.Println("Pop:", minH.Pop())

	fmt.Println("\n=== Max Heap ===")
	maxH := &dsheap.MaxHeap{}
	maxH.Insert(10)
	maxH.Insert(3)
	maxH.Insert(15)
	maxH.Insert(7)
	fmt.Println("Heap:", maxH.Data)
	fmt.Println("Pop:", maxH.Pop())

	fmt.Println("\n=== Priority Queue ===")
	pq := &dsheap.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &dsheap.Item{Value: "task1", Priority: 2})
	heap.Push(pq, &dsheap.Item{Value: "task2", Priority: 5})
	heap.Push(pq, &dsheap.Item{Value: "task3", Priority: 1})

	item := heap.Pop(pq).(*dsheap.Item)
	fmt.Println("Pop Highest Priority:", item)
}
