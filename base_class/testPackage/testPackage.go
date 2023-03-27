package testPackage

import (
	"fmt"
	"strconv"
)

// golang的公有私有变量和方法 区分方式大写代表java中的public 小写代表java中的private
var A string = "testPackage变量"

func FunTest() {

	fmt.Println("进入testPackage")

	arr := [...]string{"123", "456"}

	// 遍历数组的一种方式
	for i := 0; i < len(arr); i++ {
		fmt.Println("数组元素" + arr[i])
	}

	// 使用range
	for i := range arr {
		fmt.Println("range遍历数组++", i)
	}

	//创建map的方式
	m := map[string]string{}
	m["name"] = "小明"

	m2 := make(map[string]string)
	m2["name"] = "小明"

	// 循环map的方式
	for k, v := range m2 {
		fmt.Println("k=="+k, "v=="+v)
	}

	fmt.Println("123")

}

/*
*
简单函数

接收参数和java 不一样 类型在参数后
返回值定义在收参后
return 时可以省略返回值的声明
*
*/
func MyFun(param string, arr ...string) (ret1 string, ret2 int) {

	ret1 = param

	a, _ := strconv.Atoi(param)
	ret2 = a

	return
}

func MyHandTest(param string) (reslut string) {

	//声明指针
	var a *string
	//取param的内存地址赋值给a指向
	a = &param
	// 之后给a赋值
	*a = "asd"
	fmt.Println(*a)
	return param
}

func MyHandArrTest(param []string) (reslut []string) {

	num := 0
	for i := 0; i < len(param); i++ {
		num++
		param[i] = strconv.Itoa(num)
	}

	return param
}

func MyIdeaTest(param []string) (reslut map[string]string) {

	m := make(map[string]string)

	for k, y := range param {
		key := strconv.Itoa(k)
		m[key] = y
	}
	m["arrlen"] = strconv.Itoa(len(param))
	return m
}

// 声明一个接口
type Myinter interface {
	Eat()

	Run()
}

// 声明用户实体类
type User struct {
	Name string
	Age  int
}

// go中实现接口 使用Run() 函数表示 与java中的import 关键字同理
// 之后重写interface中的方法就可以了
func (u User) Run() {

	fmt.Println(u.Name + "开始跑了！！run")
}

func (u User) Eat() {

	fmt.Println(u.Name + "开始吃饭了！！eat")
}

// 练习channl 使用
func ChannlTest() {

	// 创建一个有1个缓冲区的channle
	// ch := make(chan int, 1)
	// ch <- 1
	// fmt.Println(<-ch)

	// 无缓冲区的channl
	// 如果想要放元素 在一个线程中同时进行的时候会出现死锁的问题
	// 可以使用 go 协程函数来避免死锁问题
	// 他会先执行 main线程 <- (取)的操作
	// 之后发现channl中没有数据 通过go 函数执行另一个操作进入放的操作之后在取
	// ch := make(chan int)

	// go func() {
	// 	ch <- 1
	// }()

	// fmt.Println(<-ch)

	// 缓冲区只能存一部分元素时候
	// 先执行取的操作 之后 执行放的操作
	// 存满channl 之后执行取的操作
	// 取出一个元素 channl 空出一个空间之后在存 一个元素 再取元素交叉运行
	// ch := make(chan int, 5)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	// 使用 range 取所有元素的时候 必须使用close 函数关闭channl

	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

}
