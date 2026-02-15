/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    height := Left(root)
    
    return helper(root, height)
}

func Left(root *TreeNode) int {
    height := 0
    for root != nil {
        height++
        root = root.Left
    }
    return height
}

func Right(root *TreeNode) int {
    height := 0
    for root != nil {
        height++
        root = root.Right
    }
    return height
}


func helper(root *TreeNode, h int) int{
   if root == nil {
        return 0
    }
    lr := Right(root.Left)
    ll := h-1

    l, r := 0, 0
    if lr < ll {
        l = helper(root.Left, h-1)
        r = int(math.Pow(2.0, float64(h-2))) - 1
    } else {
        l = int(math.Pow(2.0, float64(h-1))) - 1
        r = helper(root.Right, h-1)
    }

    return 1 + l + r
}
