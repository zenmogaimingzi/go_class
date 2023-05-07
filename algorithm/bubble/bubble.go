package main

import "fmt"

func main() {

	//fmt.Print("123")
	arr := [8]int{2, 4, 6, 8, 1, 3, 5, 6}
	MyBubble(arr[:])
	for v := range arr {
		fmt.Print(v)
	}
}

func MyBubble(arr []int) {

	// 用于标记循环次数
	var num = 1
	// 用于标记是否发生交换
	var flag = false
	for i := 0; i < len(arr)-1; i++ {
		num++
		for j := 0; j < len(arr)-i-1; j++ {
			num++
			if arr[j] > arr[j+1] {
				// 交换arr[j]和arr[j+1]的值
				flag = true
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		if !flag {
			break
		}
	}

	fmt.Println(num)
}
