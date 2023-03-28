package tree
//二叉搜索树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 查找 二叉搜索树查找指定值
func search(root *TreeNode,Val int) *TreeNode{
	if root==nil{
		return nil
	}
	if root.Val==Val{
		return root
	}
	if root.Val>Val{
		return search(root.Left,Val)
	}
	return search(root.Right,Val)
}



//最大值 二叉搜索树查找最大值
func findMax(node *TreeNode)*TreeNode{
	if node==nil{
		return nil
	}
	for node.Right!=nil{
		node=node.Right
	}
	return node
}



//最小值

func findMin(node *TreeNode)*TreeNode{
	if node==nil{
		return nil
	}
	for node.Left!=nil{
		node=node.Left
	}
	return node
}
//父节点
func searchPearent(root *TreeNode,Val int)*TreeNode{
	if root==nil{
		return nil
	}
	if root.Left==nil&&root.Right==nil{
		return nil
	}
	if root.Left.Val==Val||root.Right.Val==Val{
		return root
	}
	if root.Val>Val{
		return searchPearent(root.Left,Val)
	}
	return searchPearent(root.Right,Val)
}
//
//
//前驱
//
func findqianqu(root *TreeNode,Val int)*TreeNode{

	node:=search(root,Val)
	if node!=nil&&node.Left!=nil{
		return findMax(node.Left)
	}
	// 两种可能，1。无前驱 2. 本身为叶子节点最大，前驱为parent
	// 无左子树，
	//      4
	//   2      5
	// 1    3

	//1 无前驱   1
	// 前驱为parent 3
	//总结找parent 且node 不是parent的左节点
	parent:=searchPearent(root,node.Val)
	for parent!=nil && parent.Left==node{
		node=parent
		parent=searchPearent(root,Val)
	}
	return parent

}
//
//后继
func findhouji(root *TreeNode,Val int )*TreeNode{

	node:=search(root,Val)
	if node!=nil&&node.Right!=nil{
		return findMin(node.Left)
	}
	// 两种可能，1。无前驱 2. 本身为叶子节点最大，前驱为parent
	// 无左子树，
	//      4
	//   2      5
	// 1    3

	//1 无前驱   1
	// 前驱为parent 3
	//总结找parent 且node 不是parent的左节点
	parent:=searchPearent(root,node.Val)
	for parent!=nil && parent.Right==node{
		node=parent
		parent=searchPearent(root,Val)
	}
	return parent

}
