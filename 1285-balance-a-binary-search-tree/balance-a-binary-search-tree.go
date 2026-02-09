/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    val := []int{}
    inOrder(root,&val)
    return build(&val,0,len(val)-1)
}


func inOrder(root *TreeNode , val *[]int) {
    if root == nil {
        return
    }

    inOrder(root.Left , val)
    *val = append(*val , root.Val)
    inOrder(root.Right , val)
}

func build(val *[]int , left , right int) (newNode *TreeNode){
    if left > right {
        return nil
    }

    mid := left + (right - left) / 2

    return &TreeNode{
        Val:   (*val)[mid],
        Left:  build(val, left, mid-1),
        Right: build(val, mid+1, right),
    }
}