package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	output := []int{}

	if root == nil {
		return nil
	}

	output = append(output, inorderTraversal(root.Left)...)
	output = append(output, root.Val)
	output = append(output, inorderTraversal(root.Right)...)

	return output
}

func main() {

}
