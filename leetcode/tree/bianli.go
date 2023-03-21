package tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 二叉树的前序遍历
// 根在前
// 以根左右 方式进行遍历
var res1 []int
func qianxu(root *TreeNode)[]int {
	res1=make([]int,0)
	dfsqian(root)
	return res1
}

func dfsqian(root *TreeNode){
	if root==nil{
		return
	}
	res1 = append(res1,root.Val)
	dfsqian(root.Left)
	dfsqian(root.Right)
}
// 方法二
func qianxu2(root *TreeNode)[]int {
	res1:=make([]int,0)

	if root==nil{
		return nil
	}
	// 棧
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	for len(stack)!=0{
		top:=stack[len(stack)-1]
		res1 =append(res1,top.Val)
		if top.Right!=nil{
			stack=append(stack,top.Right)
		}
		if top.Left!=nil{
			stack=append(stack,top.Left)
		}
	}


	return res1
}






// 二叉树的中序遍历
// 根在中
func zhongxu(root *TreeNode)[]int {
	res1=make([]int,0)
	dfszhong(root)
	return res1
}

func dfszhong(root *TreeNode){
	if root==nil{
		return
	}
	dfsqian(root.Left)
	res1 = append(res1,root.Val)
	dfsqian(root.Right)
}
// 方法二
func zhongxu2(root *TreeNode)[]int {
	res1:=make([]int,0)

	if root==nil{
		return nil
	}
	// 棧
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	visited:=make(map[*TreeNode]bool)
	for len(stack)!=0{


		top:=stack[len(stack)-1]
		visit,ok:=visited[top]
		if top.Left!=nil&&(!visit||!ok){
			stack=append(stack,top.Left)
		}else{

			res1=append(res1,top.Val)
			visited[top]=true
			stack= stack[:len(stack)-1]
			if top.Right!=nil{
				stack=append(stack,top.Right)
			}

		}
	}
	return res1
}

// 二叉树后序遍历
// 根在后
func houxu(root *TreeNode)[]int {
	res1=make([]int,0)
	dfshou(root)
	return res1
}

func dfshou(root *TreeNode){
	if root==nil{
		return
	}
	dfsqian(root.Left)
	dfsqian(root.Right)
	res1 = append(res1,root.Val)

}

func houxu2(root *TreeNode)[]int{
	res1:=make([]int,0)
	if root==nil{
		return res1
	}
	visited:=make(map[*TreeNode]bool)
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	for len(stack)!=0{
		top:=stack[len(stack)-1]
		visit,ok:=visited[top]
		if top.Right!=nil&&(!visit||!ok){
			stack=append(stack,top.Right)
		}else{

			res1=append(res1,top.Val)
			visited[top]=true
			stack= stack[:len(stack)-1]
			if top.Left!=nil{
				stack=append(stack,top.Left)
			}

		}

	}

	return res1


}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func postorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}

	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	visited:=make(map[*TreeNode]bool)
	for len(stack)!=0{
		top:=stack[len(stack)-1]
		visitbool,ok:=visited[top.Left]
		if top.Left!=nil && (!visitbool||!ok){
			stack=append(stack,top.Left)
		}else{
			res = append(res,top.Val)
			stack = stack[:len(stack)-1]
			visited[top]=true

			if top.Right!=nil{
				stack=append(stack,top.Right)
			}
		}

	}
	return res
}