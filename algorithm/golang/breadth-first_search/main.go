package main

import "fmt"

//BFS to print the tree in breadth first fashion
type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Repeat int
	Parent *Tree
}

//Insert method to insert new values in the trees
//Sice every node is itself a Tree struct, we can use
//Tree struct of each particular node
func (t *Tree) Insert(v int) error {

	switch {

	case v == t.Value:
		t.Repeat++

	//if value is less than the value at the node and

	//if Left is empty, add a node here only.
	//if left is not empty, recrusion
	case v < t.Value:

		if t.Left == nil {
			t.Left = &Tree{v, nil, nil, 1, t}
			return nil
		}

		return t.Left.Insert(v)

	//if value is greater than the value at the node and

	//if right is empty, add a node here only.
	//if right is not empty, recrusion
	case v > t.Value:

		if t.Right == nil {
			t.Right = &Tree{v, nil, nil, 1, t}
			return nil
		}

		return t.Right.Insert(v)

	}

	return nil

}

//Traversal Interfcae to implement
type Traversal interface {
	Initialize() []int
	PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int
	util(stack []*Tree, res []int, visited map[int]bool) ([]int, []*Tree)
}

type defImplement struct {
}

func (*defImplement) Initialize() []int {
	return []int{}
}

func (*defImplement) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	return []int{}
}

func (*defImplement) util(stack []*Tree, res []int, visited map[int]bool) ([]int, []*Tree) {
	visited[stack[len(stack)-1].Value] = true
	res = append(res, stack[len(stack)-1].Value)
	stack = stack[:len(stack)-1]
	return res, stack
}

//Inorder DFS with InOrder, embedding Traversal interface
// we will intialize this struct with defImplement as Traversal
//type to  inherit util implementation
type Inorder struct {
	tree *Tree
	Traversal
}

func (ord *Inorder) Initialize() []int {
	stack := []*Tree{}
	stack = append(stack, ord.tree)
	result := []int{}
	visited := map[int]bool{}
	return ord.PrintTraversal(stack, result, visited)

}

func (ord *Inorder) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:
		stack = append(stack, subTree.Left)
		return ord.PrintTraversal(stack, res, visited)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:

		res, stack = ord.util(stack, res, visited)
		stack = append(stack, subTree.Right)
		return ord.PrintTraversal(stack, res, visited)

	default:
		res, stack = ord.util(stack, res, visited)
		return ord.PrintTraversal(stack, res, visited)

	}

}

type Postorder struct {
	tree *Tree
	Traversal
}

func (ord *Postorder) Initialize() []int {
	stack := []*Tree{}
	stack = append(stack, ord.tree)
	result := []int{}
	visited := map[int]bool{}
	return ord.PrintTraversal(stack, result, visited)

}

//PrintTraversal to print traversal of the binary tree
func (ord *Postorder) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:
		stack = append(stack, subTree.Left)
		return ord.PrintTraversal(stack, res, visited)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:

		stack = append(stack, subTree.Right)
		return ord.PrintTraversal(stack, res, visited)

	default:
		res, stack = ord.util(stack, res, visited)
		return ord.PrintTraversal(stack, res, visited)

	}

}

type Preorder struct {
	tree *Tree
	Traversal
}

func (ord *Preorder) Initialize() []int {
	stack := []*Tree{}
	stack = append(stack, ord.tree)
	result := []int{ord.tree.Value}
	visited := map[int]bool{
		ord.tree.Value: true,
	}

	return ord.PrintTraversal(stack, result, visited)

}

func (ord *Preorder) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:

		res = append(res, subTree.Left.Value)
		visited[subTree.Left.Value] = true
		stack = append(stack, subTree.Left)
		return ord.PrintTraversal(stack, res, visited)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:
		res = append(res, subTree.Right.Value)
		visited[subTree.Right.Value] = true
		stack = append(stack, subTree.Right)
		return ord.PrintTraversal(stack, res, visited)

	default:
		stack = stack[:len(stack)-1]
		return ord.PrintTraversal(stack, res, visited)

	}

}

// BinaryTree struct which actually makes a binary tree out of the integer array called as numbers
func BinaryTree(numbers []int) *Tree {
	root := numbers[0]
	tree := Tree{root, nil, nil, 0, nil}
	for _, value := range numbers[1:] {
		tree.Insert(value)
	}
	return &tree
}

func main() {
	numbers := []int{53, 23, 19, 5776, 170, 223, 45, 75, 90, 802, 63, 29, 3, 887, 456, 24, 2, 21, 34, 49, 6555}

	tree := BinaryTree(numbers)

	f := Inorder{tree, &defImplement{}}
	fmt.Printf("Inorder Tree traversal %v\n", f.Initialize())

	post := Postorder{tree, &defImplement{}}
	fmt.Printf("Post Tree traversal %v\n", post.Initialize())

	pre := Preorder{tree, &defImplement{}}
	fmt.Printf("Pre Tree traversal %v\n", pre.Initialize())

}
