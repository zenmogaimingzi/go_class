package brack

import "fmt"

func MyBrack() {

	// 跳转语句练习
A:
	for i := 0; i < 10; i++ {
		if i > 3 {
			// 跳出A循环
			break A
			// 进入B循环体
			goto B
		}
	}

B:
	fmt.Println("进入B方法")

}
