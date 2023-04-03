###     矩阵问题

动态规划 DFS 深度遍历计算
```
1.写深度递归方式（自然思想）
2.dfs 转为记忆化搜素（增加dp 空间）
3.根据dfs 的边界条件判断特殊值，根据特殊值优化代码 （486）


```
回溯问题
增量构造答案过程，， 取球等问题
注意边界后改成dfs


思考三个问题  
1.当前单步操作      （构造第i步） 仿照动态规划，入参i 为大于i的部分   
2.子问题是什么       构造下标>i场景  
2下一个子问题是什么   构造下标>i+1场景  



两种模板
1.从输入思考
2.从输出思考




###     链表

基本循环遍历操作

```
for list!=nil{
	list =list.Next
}
```

定义初始节点和保护节点

```
 node :=&List{Val:0,Next:nil}
 project:=node

```

快慢指针操作（取链表中间 判断链表是否循环）

```
 slow,fast:=head,head
 for fast!=nil&&fast.Next!=nil&& slow!=nil{
   fast.Next =fast.Next.Next
   slow = slow.Next
 
 }

```

###     数组问题

求最短长度类似题目：双指针 求最长长度类似题目：双指针 or 滑动窗口 求比当前数字大的数字list 单调栈

数组仅两个数值操作时候，考虑使用hash处理

双指针模板 数组中指定值

```go
l, r := 0, len(s)-1
for l<r{
l++
r--
}


```

滑动窗口模板

最小数组问题

```go
l, r := 0, 0
for r<len(s){
r++
for {
l++
}
}
```

滑动窗口+前缀和

```go
//数组全为正 或者相关数组连续 或者除法过程中前缀和单调递增
//求前缀和同时进行滑动窗口处理   209 713
// 如果数据有正负 或者数据期间增加+减小 前缀和+归并 327
func minSubArrayLen(target int, nums []int) int {
    res:=math.MaxInt64
    sum:=0
    l,r:=0,0
     for r<=len(nums)-1{
         sum+=nums[r]
         for sum >=target{
            res=min(res,r-l+1)
            sum-=nums[l]
            l++
        }
            r++
        }
     if res==math.MaxInt64{
     res=0
     }
	 return res
}



```

单调栈模板

```go

package test

//求list 中下一个比当前数字大的数字
//tmp 为栈 tmp 可能是单调递减的（求最大） 也可能是单调递增的（面积 87） 取最小高
func A(listA []int, target int) []int {
	tmp := make([]int, 0)
	res := make([]int, len(listA))
	for k, _ := range listA {
		for len(tmp) > 0 && k > listA[tmp[len(tmp)-1]] {
			index := tmp[len(tmp)-1]
			res[index] = k - index
			tmp = tmp[:len(tmp)-1]

		}
		tmp = append(tmp, k)
	}
	return res
}

```

前缀和模板
注意还有前缀亦或和 前缀余（同余公式）
```
package test

func A(listA []int) int {
	// 求连续子数组和 
	// 求连续子数组中奇数、偶数个数
	nums := make([]int, len(listA)+1)
	for k, v := range listA {
		nums[k+1] = nums[k] + v
	}

	return nums[0]

}



```

###     链表问题

需要熟悉链表操作 1.反转链表

```
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode

    node := head
    for node!=nil{
        next:=node.Next
        node.Next = pre
 
        pre = node
        node =next
    }

    return pre
}

```

###     二叉树

dfs 前序中序后续遍历

二叉搜素树

###     递归

递归场景 回溯算法 78

###     位图运算
 go 几种位图运算  
 & 与     都是1则为1   与自己都是自己，与0是0  与intmax 是intmax   a&1 相当于取最后1位
 | 或     有1个1 则为1  或自己是自己  或0 是自己 或intmax 是intmax  a|1 相当于将最后1位置1 -> 指定位置置为1（a|=1<<n） 
 ^ 亦或    都是1 则为0    亦或自己是0   
>> 双目  右移动n位  >>1 相当于/2的1次方
<< 双目  左移动n位  <<1 相当于*2的1次方  


常见位图逻辑
```go
//1.创建一个int数字
//2.

// 取每一位
for i:= 0; i < 32; i++ {
    tmp:= int32(num) >> i & 1
}



```

###     公式计算

1.快速幂算法、矩阵快速幂算法 m的k次方 k 转为二进制 根据二进制进行计算 如果是1 则a=a^k*2  (k=0,1,2...)

````azure

````

2.大量幂函数 分治 将大的幂拆成较小的进行计算 可以使用递归