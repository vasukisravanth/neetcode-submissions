/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {

	if root==nil{
		return 0
	}

	queue:=[]*TreeNode{root}
	height:=0

	for len(queue)>0{
		levelsize:=len(queue)
		for i:=0;i<levelsize;i++{
			node:=queue[0]
			queue=queue[1:]

			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}
		height++

	}

	return height
    
}
