package tree


type TreeNode struct {
    Val   int
   Left  *TreeNode
   Right *TreeNode
}

// 验证 二叉搜素树的最近公共父节点
// 递归 +分类
// 给两个节点 p q 判断pq的公共父节点
// 二叉搜素树 左边的都小于父节点 右边的都大于父节点
// 分类讨论
// p q 都大于 根节点 则递归右侧
// pq 都小于根节点则递归左侧
// pq 分别在两边 则root 就是根
func searchfather(root *TreeNode ,p,q *TreeNode) *TreeNode{
	x:=root
	if p.Val>x.Val && q.Val>x.Val{

		searchfather(root.Left,p,q)

	}
	if p.Val<x.Val && q.Val<x.Val{

		searchfather(root.Right,p,q)

	}

	return root
}
