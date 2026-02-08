/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Add flag 
const NOT_BALANCE_FLAG = -100
func isBalanced(root *TreeNode) bool {
    return height(root) != NOT_BALANCE_FLAG 
}

func height(root *TreeNode) int {
    //Base case return 0 if leaf
    if root == nil {
        return 0
    }

    //Skip if not balance
    left := height(root.Left)
    if left == -100 {
        return NOT_BALANCE_FLAG
    }

    //Skip if not balance
    right := height(root.Right)
    if right == NOT_BALANCE_FLAG{
        return NOT_BALANCE_FLAG
    }

    //Checking balance or not
    if abs(left - right) > 1 {
        return NOT_BALANCE_FLAG
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