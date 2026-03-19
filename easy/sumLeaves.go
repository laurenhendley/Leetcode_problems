/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root, false)
}

func helper(root *TreeNode, check bool) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil && check {
        return root.Val
    }
    
    left := helper(root.Left, true)
    right := helper(root.Right, false)

    return left + right
}
