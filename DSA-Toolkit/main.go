package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// Sortings
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
