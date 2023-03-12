package dp

import (
	"fmt"
	"testing"
)

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
//注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。


func wordBreak(s string, wordDict []string) bool {
	tmpdict:=make(map[string]bool)
	for _,v:=range wordDict{
		tmpdict[v]=true
	}
	dp:=make([]bool,len(s)+1)
	dp[0]=true
	for i:=1;i<=len(s);i++{
		for j:=0;j<i;j++{
			if dp[j]&&tmpdict[s[j:i]]{
				dp[i]=true
				break
			}
		}
	}



	return dp[len(s)]
}




//func dfs(s string,tmp string,dict []string) bool{
//	if s==tmp{
//		return true
//	}
//	if len(tmp)>=len(s){
//		return false
//	}
//
//	// aab    aa
//	for i:=0;i<=len(dict)-1;i++{
//		fmt.Println(tmp+dict[i])
//		flag := dfs(s,tmp+dict[i],dict,)
//		if flag{
//			return true
//		}
//	}
//
//	return  false
//}


func TestWordBreak(t *testing.T){
	dict:=[]string{"leet","code"}
	str:="leetcode"
	res:= wordBreak(str,dict)
	fmt.Println(res)
	//a:=make(map[string]bool)
	a:=map[string]string{
		"a":"a",}
	fmt.Println(a["b"])

}