
type TreeNode struct {
    Val int
	Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if depthCheck(root) == -1 {
        return false
    }
    return true
}

func depthCheck(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh := depthCheck(root.Left)
    if lh == -1 {
        return -1
    }

    rh := depthCheck(root.Right)
    if rh == -1 {
        return -1
    }

    if abs(rh - lh) > 1{
        return -1
    }

    return 1 + max(lh, rh)
}

func abs(x int) int{
    if(x < 0){
        return -x
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}