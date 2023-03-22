package tree

//给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//
//树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
type Node struct {
     Val int
    Children []*Node
}


func levelOrder(root *Node) [][]int {
	if root ==nil{
		return nil
	}

	queue:=make([]*Node,0)
	res:=make([][]int,0)
	queue=append(queue,root)
	for len(queue)!=0{
		tmp:=make([]int,0)
		size:=len(queue)
		for i:=0;i<size;i++{
			top:=queue[0]
			tmp=append(tmp,top.Val)
			queue=queue[1:]
			queue=append(queue,top.Children...)
		}
		res=append(res,tmp)

	}

	return res


}
