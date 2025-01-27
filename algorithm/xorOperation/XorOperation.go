/*
 * @Description:
 * @Version: 2.0
 * @Autor: lhl
 * @Date: 2023-05-13 18:26:24
 * @LastEditors: lhl
 * @LastEditTime: 2023-05-14 16:57:57
 */
package main

import (
	"fmt"
)

func main() {

	// num := 42
	// fmt.Printf("二进制表示：%b\n", num)
	// fmt.Printf("二进制表示：%b\n", ^num)

	// num21 := uint(42)
	// // 对二进制数取反
	// reversedNum := num21 + 1 ^ uint(^uint(0))

	// // fmt
	// fmt.Printf("取反后的二进制表示：%b\n", reversedNum)

	arr := [9]int{2, 2, 2, 3, 3, 4, 4, 5, 5}
	i := TakeOddNumber(arr[:])
	fmt.Println("取数组中奇数个的值：", i)

	arr2 := [10]int{2, 2, 5, 3, 3, 4, 4, 4, 5, 5}
	m := TwoTakeOddNumber(arr2[:])
	for key, value := range m {
		fmt.Println("取数组中奇数个的值：", value, key)
	}

}

// TakeOddNumber  通过^(异或运算查找数组中的奇数个的整数) 时间复杂度BigO(1)
//
//	@param arr
//	@return num
func TakeOddNumber(arr []int) (num int) {
	// 通过range便利集合时候注意索引的存在如果只用到索引就可以使用一个参数接受range的返回值，
	// 如果需要对元素进行操作那需要返回两个参数第一个代表索引，第二个代表内容
	for _, item := range arr {
		num ^= item
	}
	return num
}

// TwoTakeOddNumber 通过异或运算查找两个奇数的整数 时间复杂度BigO(1)
//
//	@param arr
//	@return num
func TwoTakeOddNumber(arr []int) (reulstArr map[int]int) {

	// 异或算出其中一个奇数个的整数
	num := 0
	for _, item := range arr {
		num ^= item
	}
	//
	rightOne := num & (^num + 1)

	fmt.Printf("二进制表示：%b\n", ^rightOne)
	for _, item := range arr {
		if item&rightOne == 1 {
			rightOne ^= item
		}

	}
	num3 := num ^ rightOne

	m := make(map[int]int)
	m[1] = num3
	m[2] = rightOne
	return m
}
