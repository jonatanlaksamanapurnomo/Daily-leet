/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return height(root) != -100
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := height(root.Left)
    if left == -100 {
        return -100
    }

    right := height(root.Right)
    if right == -100{
        return -100
    }

    //Checking balance or not
    if abs(left - right) > 1 {
        return -100
    }


    //Get Height Tree
    return max(left,right) + 1
}

func max(a,b int ) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}