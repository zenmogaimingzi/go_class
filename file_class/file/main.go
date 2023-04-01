package main

import (
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
			return
		}

		fmt.Println(string(b), n)
	}

}
