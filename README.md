# GolangDocs

# 入门指南
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


