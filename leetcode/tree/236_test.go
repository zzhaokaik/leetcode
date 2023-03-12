package tree

// 验证 二叉树的最近公共父节点
// 递归 +分类
// 给两个节点 p q 判断pq的公共父节点
// 分类讨论 递归节点分别讨论
//	root =nil 返回 nil // 其实就是root

//	root 为p q 直接返回
//  left right 有1个为nil 返回另一个    pq 在同一分支 或者对应分支没有pq 返回另一个分支
// 左右都不为nil 返回root 说明左右各有pq 则一定是公共父节点 返回root pq 在不同分支



func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil || root==p || root==q{
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left!=nil && right!=nil{
		return root
	}
	if left!=nil{
		return left
	}
	return right
}
