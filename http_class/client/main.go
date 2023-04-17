package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {

	// 测试GET请求
	// c := new(http.Client)
	// requset, _ := http.NewRequest("GET", "http://127.0.0.1:8080/test", nil)
	// res, _ := c.Do(requset)
	// body := res.Body
	// b, _ := io.ReadAll(body)
	// fmt.Print(string(b))

	// 测试POST请求
	postClient := new(http.Client)

	posRequset, _ := http.NewRequest("POST", "http://127.0.0.1:8080/test", bytes.NewBuffer([]byte("{\"test\":\"客服端\"}")))
	postRes, _ := postClient.Do(posRequset)
	postBody := postRes.Body
	postBf, _ := io.ReadAll(postBody)
	fmt.Print(string(postBf))

}
