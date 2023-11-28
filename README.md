# GolangDocs

# 入门指南

go run <filename>

package command-line-arguments is not a main package

问题产生原因：

文件开头应该是：package main，golang 和 c 一样以 main 作为入口函数，且包名需要是 mian
```
package main
import "fmt"
func main(){
    fmt.Println("hello world.")
}
```

