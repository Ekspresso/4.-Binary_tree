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
	arr := dfsInOrder(tree)
	fmt.Println("inOrder: ", arr)
	arr = dfsPostOrder(tree)
	fmt.Println("postOrder: ", arr)
	arr = dfsPreOrder(tree)
	fmt.Println("preOrder: ", arr)
}

func dfsInOrder(root *BinaryTree) []int {
	if root != nil {
		var arr []int
		arr = append(arr, dfsInOrder(root.Left)...)
		arr = append(arr, root.Val)
		arr = append(arr, dfsInOrder(root.Right)...)
		return arr
	}
	return nil
}

func dfsPostOrder(root *BinaryTree) []int {
	if root != nil {
		var arr []int
		arr = append(arr, dfsPostOrder(root.Left)...)
		arr = append(arr, dfsPostOrder(root.Right)...)
		arr = append(arr, root.Val)
		return arr
	}
	return nil
}

func dfsPreOrder(root *BinaryTree) []int {
	if root != nil {
		var arr []int
		arr = append(arr, root.Val)
		arr = append(arr, dfsPreOrder(root.Left)...)
		arr = append(arr, dfsPreOrder(root.Right)...)
		return arr
	}
	return nil
}
