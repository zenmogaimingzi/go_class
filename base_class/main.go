package main

import (
	"fmt"
	testPackage "go_class/base_class/testPackage"
	"sync"
)

func main() {
	// var str string = "你好 世界！！"
	// //隐式声明变量
	// b := str
	// fmt.Println(testPackage.A)
	// fmt.Println(b)

	// var strint string = "123"

	// convInt, _ := strconv.Atoi(strint)

	// int64, _ := strconv.ParseInt(strint, 10, 64)

	// fmt.Printf("%T", int64)

	// fmt.Printf("%T", convInt)

	// if convInt == 123 {
	// 	fmt.Println("convInt==", convInt)
	// } else {
	// 	fmt.Println("走了else!!")
	// }

	b := 10
	switch b {
	case 0:
		fmt.Println("b==0")
	case 10:
		fmt.Println("b==10")
	case 12:
		fmt.Println("b==12")
	default:
		fmt.Println("都不是走了默认逻辑！")
	}

	// for 循环的三种写法

	// 第一种写法
	num := 0
	for {
		b++
		num++
		if b > 20 {
			fmt.Println("b大于20")
			break
		}
	}

	fmt.Println("循环次数", num)

	//第二种写法
	for a := 0; a < 3; a++ {
		fmt.Println(a)
	}

	//第三种方式
	c := 0
	for c < 3 {
		c++
		fmt.Println("第三种方式", c)
	}

	testPackage.FunTest()

	arr := []string{"1", "2", "3"}

	ret1, ret2 := testPackage.MyFun("123", arr...)

	fmt.Println("返回值ret1", ret1, "===返回值ret2", ret2)

	s := testPackage.MyHandTest("233")
	fmt.Println(s)

	ff := testPackage.MyHandArrTest([]string{"9", "8", "7"})

	fmt.Println(ff)

	gg := testPackage.MyIdeaTest([]string{"9", "8", "7"})

	fmt.Println(gg)

	// 学习interface

	// var myInter testPackage.Myinter

	// uu := testPackage.User{
	// 	Name: "小明",
	// 	Age:  12,
	// }
	// myInter = uu

	// 直接将User实体类交给 myInter接口
	var myInter testPackage.Myinter

	myInter = testPackage.User{
		Name: "小明",
		Age:  12,
	}

	myInter.Run()
	myInter.Eat()

	// go 函数
	//线程计数器 waitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go Threa(&wg)
	wg.Wait()

	// 练习go channle
	testPackage.ChannlTest()

}

func Threa(wg *sync.WaitGroup) {
	fmt.Println(" 调用 Thread！！！")
	wg.Done()
}
