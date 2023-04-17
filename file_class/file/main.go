package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.OpenFile("D:/work/open_go_projects/go_class/file_class/resources/123.txt", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		b := make([]byte, 12)
		n, err := f.Read(b)

		if err != nil {
			fmt.Println(err)
			// 使用goto 调用写的操作
			goto B
			return
		}

		fmt.Println(string(b), n)
	}

B:
	// 此方法直接写入并覆盖之前的内容
	//f.WriteString("写入内容！！")

	// 此方法不会覆盖之前内容只会追加到内容最后
	// 此方法创建一个缓存的写操作
	w := bufio.NewWriter(f)
	// 将内容放到bufio缓存中
	w.WriteString("使用bufio写入内容")
	// 刷新文件bufio内容生效 写入文件
	w.Flush()

}
