package leetcode


//滑动窗口
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。


// 最大最小问题
func lengthOfLongestSubstring(s string) int {
	i,j:=0,0
	tmpmap:=make(map[byte]int)
	res:=0
	for j<len(s){
		str1:=s[j]
		tmpmap[str1]++
		j++
		for tmpmap[str1] >1{
			str2:=s[i]
			if _,ok:=tmpmap[str2];ok{
				tmpmap[str2]--
			}
			i++
		}
		res =max(res,j-i)

	}
	return res
}

func max3(a,b int)int{
	if a>b{
		return a
	}
	return b
}
