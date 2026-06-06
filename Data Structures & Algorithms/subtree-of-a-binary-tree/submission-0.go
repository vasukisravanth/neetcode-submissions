/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsSameTree(a *TreeNode,b *TreeNode)bool{
    if a==nil && b==nil{
        return true
    }
    if a==nil || b==nil{
        return false
    }

    if a.Val!=b.Val{
        return false
    }

    return IsSameTree(a.Left,b.Left)&&IsSameTree(a.Right,b.Right)
    

}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root==nil{
        return false
    }
 
    if IsSameTree(root, subRoot) {
        return true
    }

    return isSubtree(root.Left, subRoot) ||
           isSubtree(root.Right, subRoot)

    
}
