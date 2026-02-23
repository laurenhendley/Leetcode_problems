/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "strings"
    "strconv"
)

func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }
    return dfs([]string{}, root)
}

func dfs(res []string, root *TreeNode) []string {
    if root == nil {
        return nil
    }

    res = append(res, strconv.Itoa(root.Val))
    if root.Left == nil && root.Right == nil{
        return []string{strings.Join(res, "->")}
    }

    l := dfs(res,root.Left)
    r := dfs(res,root.Right)

    return append(l, r...)
}
