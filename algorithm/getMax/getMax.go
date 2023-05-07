package main

import "fmt"

func main() {

	arr := []int{1, 3, 5, 2, 4, 6}
	result := MyProcess(arr, 0, len(arr)-1)
	// Output: 6
	fmt.Println(result)

}

// 时间复杂度为 O(N)
// 使用二分取中法取最大值
func MyProcess(arr []int, L int, R int) (reslut int) {

	if L == R {
		return arr[L]
	}
	// 取中点
	// 使用 L+R/2 在数据过大的情况时会出现异常
	// 改为这个操作改为 L+((R-l)/2)
	// R-L/2 可以简化为 二字节右移一位 改为 (R - L) >> 1
	// 最后得出 中点 :=  L + ((R - L) >> 1)
	mid := L + ((R - L) >> 1)
	leftMax := MyProcess(arr, L, mid)
	rightMax := MyProcess(arr, mid+1, R)

	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
