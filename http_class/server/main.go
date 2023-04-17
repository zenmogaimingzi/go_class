package main

import (
	"fmt"
	"io"
	"net/http"
)

func hander(res http.ResponseWriter, req *http.Request) {

	m := req.Method

	switch m {
	case "GET":
		fmt.Println("GET请求")
		res.Write([]byte("接受GET请求，返回数据"))
		break
	case "POST":
		fmt.Println("POST请求")

		body, _ := io.ReadAll(req.Body)
		// 配合对象将json字符串转换为对象 json.Unmarshal(jsonStr,Object)
		// http包定义了http状态码直接使用
		head := res.Header()
		head["test"] = []string{"str", "str2"}
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(body))
		break
	default:
		fmt.Print("未识别请求！！")
		res.Write([]byte("未识别请求！！"))
		break
	}

}

func main() {

	//http.HandleFunc("/test", hander)
	//http.ListenAndServe(":8080", nil)

	// 直接使用mux
	mux := http.NewServeMux()
	mux.HandleFunc("/test", hander)
	// 声明8080端口时候需要加:
	http.ListenAndServe(":8080", mux)

	fmt.Println("123")

}
