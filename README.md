# Golang Notes
[Golang 总会被不经意间的细节所打败](./markdown/golang.md)

[专栏 数据库和 ORM](./markdown/database.md)

[专栏 Go和Python的对比分析](./markdown/golangVSpython.md)
# GolangDocs
网站:
1. gin <https://github.com/gin-gonic/gin>
2. beego <https://github.com/beego/beego>

微服务:
1. go kit
2. Istio

容器编排
1. kubernetes
2. swarm

中间件
1. nsg
2. zinx
3. leaf
4. gRPC
5. redis

爬虫
1. goquery

https://www.yuque.com/aceld/golang/uh0124

https://zerotomastery.io/blog/golang-practice-projects

# 快速入门
go run gofile.go

package command-line-arguments is not a main package

文件开头应该是：package main，golang 和 c 一样以 main 作为入口函数，且包名需要是 mian
```
package main
import "fmt"
func main(){
    fmt.Println("hello world.")
}
```

注：go build gofile.go，会生成可执行程序，通常调试情况下我们使用 go run.

go run --work main.go 可以查看临时文件的位置

## 引号
在GO中，单引号 '' 表示的是 ASCII ， 双引号 "" 表示的是字符串， 反引号 `` 表示的是原生字符串 

gopath

go mod init

