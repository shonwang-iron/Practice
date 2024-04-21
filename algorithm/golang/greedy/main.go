package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	coins := []int{25, 10, 5, 1} // Coins with denominations of 25, 10, 5, 1
	amount := 47                 // Change amount is 47 cents

	change := makeChange(coins, amount)
	fmt.Println("The change amount is:", amount, "cents")
	fmt.Println("Change result:", change)

	fmt.Println("================================================================")
	tasks := []Task{
		{Deadline: 2, Profit: 20},
		{Deadline: 1, Profit: 10},
		{Deadline: 3, Profit: 15},
		{Deadline: 2, Profit: 5},
	}

	maxProfit := maximizeProfit(tasks)
	fmt.Println("The maximum total profit is:", maxProfit)

	fmt.Println("================================================================")
	s := "hello world"
	codeTable := HuffmanEncoding(s)

	fmt.Println("Character Huffman encoding:")
	for char, code := range codeTable {
		fmt.Printf("%c: %s\n", char, code)
	}

	fmt.Println("================================================================")
	graph := Graph{
		V: 5,
		Edges: [][]int{
			{0, 2, 0, 6, 0},
			{2, 0, 3, 8, 5},
			{0, 3, 0, 0, 7},
			{6, 8, 0, 0, 9},
			{0, 5, 7, 9, 0},
		},
	}

	PrimMST(graph)
}

func makeChange(coins []int, amount int) map[int]int {
	change := make(map[int]int)

	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		numCoins := amount / coin
		if numCoins > 0 {
			change[coin] = numCoins
			amount -= numCoins * coin
		}
	}

	return change
}

type Task struct {
	Deadline int
	Profit   int
}

// Implement the sort.Interface interface to sort Tasks
type ByDeadline []Task

func (a ByDeadline) Len() int           { return len(a) }
func (a ByDeadline) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDeadline) Less(i, j int) bool { return a[i].Deadline < a[j].Deadline }

// maximize total profit
func maximizeProfit(tasks []Task) int {
	sort.Sort(ByDeadline(tasks))

	maxDeadline := 0
	for _, task := range tasks {
		if task.Deadline > maxDeadline {
			maxDeadline = task.Deadline
		}
	}

	// Initialize a timeline and mark whether there are tasks at each time point
	schedule := make([]bool, maxDeadline)

	totalProfit := 0
	for _, task := range tasks {
		// Starting from the last time point,
		// find the earliest time when the task can be scheduled
		for i := task.Deadline - 1; i >= 0; i-- {
			if !schedule[i] {
				schedule[i] = true
				totalProfit += task.Profit
				break
			}
		}
	}

	return totalProfit
}

// Define node structure
type Node struct {
	char      rune
	frequency int
	left      *Node
	right     *Node
}

// Define priority queue for node sorting
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Huffman coding main function
func HuffmanEncoding(s string) map[rune]string {
	frequencies := make(map[rune]int)

	// Calculate character frequency
	for _, char := range s {
		frequencies[char]++
	}

	// Create a priority queue, using character frequency as priority
	pq := make(PriorityQueue, len(frequencies))
	i := 0
	for char, frequency := range frequencies {
		pq[i] = &Node{char: char, frequency: frequency}
		i++
	}
	heap.Init(&pq)

	// Build a Huffman tree
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		merged := &Node{
			char:      0,
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}

		heap.Push(&pq, merged)
	}

	// Generate Huffman encoding
	root := heap.Pop(&pq).(*Node)
	codeTable := make(map[rune]string)
	buildCode(root, "", codeTable)

	return codeTable
}

// Recursively generate Huffman coding table
func buildCode(node *Node, code string, codeTable map[rune]string) {
	if node == nil {
		return
	}

	if node.char != 0 {
		codeTable[node.char] = code
	}

	buildCode(node.left, code+"0", codeTable)
	buildCode(node.right, code+"1", codeTable)
}

// Define the adjacency matrix representation of a graph
type Graph struct {
	V     int     // number of vertices
	Edges [][]int // adjacency matrix
}

// Prim's algorithm finds the minimum spanning tree
func PrimMST(graph Graph) {
	parent := make([]int, graph.V) // Stores the parent node of the minimum spanning tree
	key := make([]int, graph.V)    // Store the minimum weight of each vertex into the minimum spanning tree

	// Initialize the key value of all vertices to infinity
	for i := range key {
		key[i] = math.MaxInt64
	}

	// Select the first vertex as the starting point
	key[0] = 0
	parent[0] = -1

	// Find V-1 vertices and add them to the minimum spanning tree
	for count := 0; count < graph.V-1; count++ {
		// Find the vertex u with the smallest key value currently
		u := minKey(key, graph.V)

		// Add vertex u to the minimum spanning tree
		for v := 0; v < graph.V; v++ {
			if graph.Edges[u][v] != 0 && graph.Edges[u][v] < key[v] {
				parent[v] = u
				key[v] = graph.Edges[u][v]
			}
		}
	}

	fmt.Println("Minimum spanning tree edges:")
	for i := 1; i < graph.V; i++ {
		fmt.Printf("%d - %d\n", parent[i], i)
	}
}

// Find the vertex with the smallest key value
func minKey(key []int, V int) int {
	min := math.MaxInt64
	minIndex := -1

	for v := 0; v < V; v++ {
		if key[v] < min {
			min = key[v]
			minIndex = v
		}
	}

	return minIndex
}
