package leetcode
//这里有 n 个航班，它们分别从 1 到 n 进行编号。
//
//有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//
//请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
//
// 
//
//示例 1：
//
//输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//输出：[10,55,45,25,25]
//解释：
//航班编号        1   2   3   4   5
//预订记录 1 ：   10  10
//预订记录 2 ：       20  20
//预订记录 3 ：       25  25  25  25
//总座位数：      10  55  45  25  25
//因此，answer = [10,55,45,25,25]



// 知识点：
// 差分数组  diff[i] = nums[i] - nums[i-1]
// 差分数组前缀和  sum[i] = diff[0] + diff[1]+...diff[i]
//               sum[i] = nums[i]
// 数组 某一区间  +n =  对应差分数组 start +n  end+1 -n
//原     1  3  5 8
//差分    1  2  2 3
//       11 13 5 8   +10
//       11  2  -8 3
//       1 3 15 18   +10
//       1  2  12 3

func corpFlightBookings(bookings [][]int, n int) []int {
	// diff 差分数组
	diff :=make([]int,n)
	for _,book:=range bookings{
		value:=book[2]
		start:=book[0]
		end:=book[1]
		diff[start-1] += value
		if end<n{
			diff[end-1+1] -=value
		}

	}
	sum:=make([]int,n+1)
	for i:=1;i<n+1;i++{
		sum[i] = sum[i-1]+diff[i-1]
	}
	return sum[1:]
}
