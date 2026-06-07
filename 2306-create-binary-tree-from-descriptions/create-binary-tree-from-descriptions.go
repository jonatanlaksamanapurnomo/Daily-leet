/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    node := make(map[int]*TreeNode)
    childMap := make(map[int]bool)

    for _,desc := range descriptions{
        par , child , isLeft := desc[0] , desc[1] , desc[2]
        childMap[child] = true
        if _,exist := node[par]; !exist{
            node[par] = &TreeNode{Val:par}
        }  
        
        if _,exist := node[child]; !exist{
            node[child] = &TreeNode{Val:child}
        }

        if isLeft == 1 {
            node[par].Left = node[child]
        } else {
            node[par].Right = node[child]
        }
    }

    for _,desc := range descriptions{
        par  := desc[0]
        _,exist := childMap[par]
        if !exist{
            return node[par]
        }
    }
    return nil
}