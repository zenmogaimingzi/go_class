## admin-demo

## 学习Golang 
### 涉及Golang基础,gin框架，Go mod依赖管理工具
#### Golang基础
##### Go 中的可见性规则
```go

//首字母大写：

//标识符（如变量、函数、类型）的首字母大写时，这个标识符是 导出（exported） 的，可以被其他包访问。
//类似于 Java 中的 public 修饰符。
func PublicFunction() {
    // 可被其他包调用
}
```
```go

//首字母小写：
//标识符的首字母小写时，这个标识符是 未导出（unexported） 的，只能在定义它的包内访问。
//类似于 Java 中的 private 修饰符。
func privateFunction() {
// 仅当前包内可调用
}

```
##### Go 中的可见性规则案例
```
项目结构
myproject/
├── main.go
└── demo/
    ├── a.go
    └── b.go

```
a.go:
```go
package demo

func TestFun() {
	helperFunction() // 可以直接调用 b.go 中的函数
}

```
b.go:
```go
package demo

func helperFunction() {
	// 一些逻辑
}

```
解释：
因为 a.go 和 b.go 都在 demo 包中，包内的所有文件可以直接访问未导出的标识符（小写开头的 helperFunction）。
这种设计避免了显式的 public 或 private 声明，只需要确保标识符是在同一个包中。

```go
package main

//  需要申明

import "myproject/demo"

func main() {
	demo.TestFun() // 可以直接调用，因为 TestFun 是导出的
}
```
a.go:
```go
package demo

// 大写开头表示导出
func TestFun() {
	helperFunction() // 可以调用包内的未导出函数
}
```

```go
package demo

// 小写开头，未导出，仅包内可见
func helperFunction() {
	// 逻辑代码
}
}
```



## 使用方式

### 打包相关
```shell
# 下载代码
git clone https://xxxx.xxx.xx/xx.git
# 更新依赖
go mod download
# 编译
go build main.go
# 清理无用mod引用
go mod tidy
# 打包启动 默认当前环境打包 windows:exe linux:可执行文件 
go build -x main.go 
#-s：省略符号表信息。
#-w：省略调试信息。
#这两个标志用于减少二进制文件的体积，提高加载和运行速度。
go build -ldflags "-s -w" main.go 
# 为 linux 生成可执行文件
GOOS=linux GOARCH=amd64 go build -o main.go 
# 如果打包提示GOOS问题使用下述命令  zipSoftware可修改打包成任意名称
$env:CGO_ENABLED="0"; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -ldflags "-s -w" -o zipSoftware main.go
# windows环境命令
$env:CGO_ENABLED="0"; $env:GOOS="windows"; $env:GOARCH="amd64"; go build -ldflags "-s -w" -o zipSoftware.exe main.go
```

### Linux中使用
```shell
#  zipSoftware是打包之后的可执行文件
# 授予可执行权限 
chmod +x zipSoftware
# liunx中使用
# './zipSoftware'是程序名称 压缩之后存放路径 压缩文件绝对路径 多个文件使用空格区分
./zipSoftware /usr/local/ /usr/local/web/redis.conf /usr/local/web/redis2.conf
```



## Thanks

<a href="https://www.jetbrains.com/?from=openwechat"><img src="https://account.jetbrains.com/static/images/jetbrains-logo-inv.svg" height="200" alt="JetBrains"/></a>



[topgoer](https://www.topgoer.com/%E5%85%B6%E4%BB%96/%E5%8E%8B%E7%BC%A9%E8%A7%A3%E5%8E%8B%E6%96%87%E4%BB%B6.html
)
##### Go 项目结构
```
项目结构
admin-demo/
├── main.go
└── demo/
    ├── a.go
    └── b.go

```
