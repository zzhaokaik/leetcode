package leetcode

import "testing"

type Node struct{
	Key int
	Value int
	Index int
	Pre *Node
	Next *Node

}

type CacheList struct{
	Head *Node
	Tail *Node
	Index int
	Size int
}

type LFUCache struct {
	MapList  map[int]*CacheList
	Capacity int
	MapNode  map[int]*Node
	LastIndex int
	Size int
}

//
// type Node struct{
//	Key int
//	Value int
//	Index int
//	Pre *Node
//	Next *Node
//
//}
//
//type CacheList struct{
//	Head *Node
//	Tail *Node
//	Index int
//	Size int
//}
//
//type LFUCache struct {
//	MapList  map[int]*CacheList
//	Capacity int
//	MapNode  map[int]*Node
//	LastIndex int
//	Size int
//}


func Constructor(capacity int) LFUCache {
	mapList:=make(map[int]*CacheList)
	mapNode:=make(map[int]*Node)
	return LFUCache{
		MapList:mapList,
		Capacity:capacity,
		MapNode:mapNode,
		LastIndex:0,
		Size:0,
	}
}


func (this *LFUCache) Get(key int) int {
	if node,ok:=this.MapNode[key];ok{

		this.remove(node)
		node.Index+=1
		this.insert(node.Index,node)

		return node.Value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if this.Capacity==0{
		return
	}
	if node,ok:=this.MapNode[key];ok{
		this.remove(node)
		node.Value = value
		this.insert(node.Index,node)
	}else{
		if len(this.MapNode)==this.Capacity{
			if minList,ok:=this.MapList[this.LastIndex];ok{
				last:=minList.Tail.Pre
				this.remove(last)
				delete(this.MapNode,last.Key)
			}
		}
		node:=&Node{
			Index:0,
			Value:value,
			Key:key,
		}
		this.insert(node.Index,node)
		this.MapNode[node.Key] = node
	}
}

func  (this *LFUCache)insert(index int,node *Node){
	var list *CacheList

	if tmp,ok:=this.MapList[index];ok{
		list = tmp
	}else{
		list=initList()
	}
	next:=list.Head.Next
	list.Head.Next = node
	node.Pre = list.Head
	node.Next = next
	next.Pre = node
	list.Size+=1
	this.MapList[index] =list
	this.MapNode[node.Key] = node
	// if len(this.MapNode)>this.Capacity{
	//     this.deleteNode()
	// }
}

func (this *LFUCache)remove(node *Node){
	if list,ok:=this.MapList[node.Index];ok{
		next:=node.Next
		node.Pre.Next = next
		next.Pre = node.Pre
		list.Size-=1
		// if list.Size==0 && this.LastIndex == node.Index{
		//     delete(this.MapList,node.Key)
		//     this.LastIndex+=1
		// }
		this.Size-=1
	}


}

// // 删除最后1个节点
// func  (this *LFUCache)deleteNode(){
//     lastIndex:=this.LastIndex
//     list,_:=this.MapList[lastIndex]

//     remonveNode:=list.Tail.Pre
//     remonveNode.Pre.Next = list.Tail
//     list.Tail.Pre = remonveNode.Pre

//     list.Size-=1
//     if list.Size==0 {
//         delete(this.MapList,this.LastIndex)
//         this.LastIndex+=1
//     }
//     delete(this.MapNode,remonveNode.Key)
//     this.Size-=1
// }


func initList() *CacheList{
	var list *CacheList
	head:=&Node{}
	tail:=&Node{}
	head.Next=tail
	tail.Pre =head
	list=&CacheList{
		Head:head,
		Tail:tail,
	}
	return list
}
/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
	mapNode:=make(map[int]*Node)
	return LFUCache{
		MapList:mapList,
		Capacity:capacity,
		MapNode:mapNode,
		LastIndex:0,
		Size:0,
	}
}


func (this *LFUCache) Get(key int) int {
	if node,ok:=this.MapNode[key];ok{

		this.remove(node)
		node.Index+=1
		this.insert(node.Index,node)

		return node.Value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if this.Capacity==0{
		return
	}
	if node,ok:=this.MapNode[key];ok{
		this.remove(node)
		node.Value = value
		this.insert(node.Index,node)
	}else{
		if len(this.MapNode)==this.Capacity{
			if minList,ok:=this.MapList[this.LastIndex];ok{
				last:=minList.Tail.Pre
				this.remove(last)
				delete(this.MapNode,last.Key)
			}
		}
		node:=&Node{
			Index:0,
			Value:value,
			Key:key,
		}
		this.insert(node.Index,node)
		this.MapNode[node.Key] = node
	}
}

func  (this *LFUCache)insert(index int,node *Node){
	var list *CacheList

	if tmp,ok:=this.MapList[index];ok{
		list = tmp
	}else{
		list=initList()
	}
	next:=list.Head.Next
	list.Head.Next = node
	node.Pre = list.Head
	node.Next = next
	next.Pre = node
	list.Size+=1
	this.MapList[index] =list
	this.MapNode[node.Key] = node
	// if len(this.MapNode)>this.Capacity{
	//     this.deleteNode()
	// }
}

func (this *LFUCache)remove(node *Node){
	if list,ok:=this.MapList[node.Index];ok{
		next:=node.Next
		node.Pre.Next = next
		next.Pre = node.Pre
		list.Size-=1
		// if list.Size==0 && this.LastIndex == node.Index{
		//     delete(this.MapList,node.Key)
		//     this.LastIndex+=1
		// }
		this.Size-=1
	}


}

// // 删除最后1个节点
// func  (this *LFUCache)deleteNode(){
//     lastIndex:=this.LastIndex
//     list,_:=this.MapList[lastIndex]

//     remonveNode:=list.Tail.Pre
//     remonveNode.Pre.Next = list.Tail
//     list.Tail.Pre = remonveNode.Pre

//     list.Size-=1
//     if list.Size==0 {
//         delete(this.MapList,this.LastIndex)
//         this.LastIndex+=1
//     }
//     delete(this.MapNode,remonveNode.Key)
//     this.Size-=1
// }


func initList() *CacheList{
	var list *CacheList
	head:=&Node{}
	tail:=&Node{}
	head.Next=tail
	tail.Pre =head
	list=&CacheList{
		Head:head,
		Tail:tail,
	}
	return list
}
/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


func TestList(t *testing.T){
	a:=Constructor(3)
	a.Put(1,1)
}