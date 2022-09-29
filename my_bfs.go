package main

import "fmt"

type BinaryTree struct {
	Val         int
	Left, Right *BinaryTree
}

func main() {
	tree := &BinaryTree{
		Val: 3,
		Left: &BinaryTree{
			Val: 1,
			Left: &BinaryTree{
				Val: 0,
			},
			Right: &BinaryTree{
				Val: 2,
			},
		},
		Right: &BinaryTree{
			Val: 5,
			Left: &BinaryTree{
				Val: 4,
			},
			Right: &BinaryTree{
				Val: 6,
			},
		},
	}
	arr := bfsOrder(tree)
	fmt.Println("bfsOrder: ", arr)
}

func bfsOrder(root *BinaryTree) []int {
	var nodes []*BinaryTree
	var arr []int
	nodes = append(nodes, root)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}
		nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		arr = append(arr, nodes[i].Val)
	}
	return arr
}
