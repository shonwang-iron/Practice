package main

import "fmt"

type Graph struct {
	nodes []*Node
}

type Node struct {
	value     int
	visited   bool
	neighbors []*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.neighbors = append(n.neighbors, neighbor)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func DFS(startNode *Node) {
	startNode.visited = true
	fmt.Printf("%v ", startNode)

	for _, neighbor := range startNode.neighbors {
		if !neighbor.visited {
			DFS(neighbor)
		}
	}
}

func main() {
	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	node5 := &Node{value: 5}

	node1.AddNeighbor(node2)
	node1.AddNeighbor(node3)
	node2.AddNeighbor(node4)
	node3.AddNeighbor(node4)
	node3.AddNeighbor(node5)

	graph := &Graph{}
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	graph.AddNode(node5)

	DFS(node1)
}
