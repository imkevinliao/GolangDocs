# golang 细节

# 指针
在C语言中，最头疼地应该就是 * 这个符号了，初学者可能觉得没什么，但是真的使用时候，在一些复杂的场景简直是炸裂。
而且对于强迫症而言，例如结构体指针的 * 符号，既可以选在结构体末尾，也可以选在变量开头，它们完全一样，但是却有两种截然不同的书写形式。

如果只是书写形式那也就算了，可怕的是指针的指针的指针，也就是使用 * 去指向一个已经是 * 的对象，然后不停套娃。

在Golang中虽然也有指针，但其实很早就听见有人说，Golang指针并不像C一样，只是一直不太清楚具体在哪里。

直到发现对于Golang，假如有一个结构体，然后声明一个指针指向它，然后当我第二次使用 * 符的时候，就产生了与C完全不同的意义，
因为第二次使用 * ，在Golang中表示的是解引用，也就是说，第二次使用的时候就还原回了结构体。完全不存在套娃行为。

忽然意识到，原来我并不是讨厌指针，我只是讨厌没有边界感的指针！

# 抽象接口
```go
type Program interface {
	hello() string
	add(int, float32) float64
	connect(string, string, int) string
}

type Python struct {
	Name   string
	Friend string
	Age    float64
	Money  float32
	Mac    int
	Hi     string
}

func (py Python) hello() string {
	return py.Hi
}

func (py Python) add(a int, b float32) float64 {
	c := float64(a) + float64(b)
	d := float64(py.Money) + py.Age
	return c + d
}

func (py Python) connect(s1, s2 string, num int) string {
	str := fmt.Sprint("%s,%s,%d", py.Name, py.Friend, py.Mac)
	return str
}
func DemoInterface() {
	python := Python{"python", "c and c++", 30, 238945723957.43545, 100, "I'm rich!"}
	fmt.Println(python.hello())
	fmt.Println(python.add(1, 1.2))
	fmt.Println(python.connect("s1", "s2", 1))
}
```
在很多语言里都有抽象方法，听了不同编程语言的说明，都是基于类进行讲解的，感觉不够清晰。而Golang这种简单明了的形式，有种醍醐灌顶的感觉。
这种将方法绑定到指定结构体的形式，真的是太奇怪了。这里的代码都是对抽象接口的实现，事实上接口的定义和实现有一定偏差也是可以编译成功的，
这点真让人困惑，毕竟对于一个连花括号是否换行都不允许的语言，竟然会有如此大的宽容度。

