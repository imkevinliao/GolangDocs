package code

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func RandomString(n int) string {
	var letters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// 为了每次生成的随机数不同，需要初始化随机数种子
	rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]int32, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomInt(min, max int) (result int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	result = min + rand.Intn(max-min+1)
	return
}

func FilePath() {
	currentDir, _ := filepath.Abs(`.`)
	fmt.Println(currentDir)
	absPath, _ := filepath.Abs("./path/to/file.txt")
	base := filepath.Base("/path/to/file.txt")
	dir := filepath.Dir("/path/to/file.txt")
	ext := filepath.Ext("/path/to/file.txt")
	path := filepath.Join("/path", "to", "file.txt")
	fmt.Println("absPath:", absPath, "\nbase:", base, "\ndir:", dir, "\next:", ext, "\npath:", path)
}

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

type MarsRover struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LaunchDate string `json:"launchDate"`
	Status     string `json:"status"`
}

func Map2Struct() {
	data := map[string]interface{}{"id": "1", "name": "Perseverance", "launchDate": "2020-07-30", "status": "Active"}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic("convert map to json failed!")
	}
	fmt.Println(string(jsonBytes))
	rover := MarsRover{}
	err = json.Unmarshal(jsonBytes, &rover)
	if err != nil {
		log.Println("convert json to struct failed!")
	}
	fmt.Println(rover)
}

func Struct2Map() {
	rover := MarsRover{1, "Perseverance", "2020-07-30", "Active"}
	jsonBytes, err := json.Marshal(rover)
	if err != nil {
		panic("convert struct to json failed!")
	}
	fmt.Println(string(jsonBytes))
	var data map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Println("convert json to map failed!")
	}
	fmt.Println(data)
}

func DemoConvert() {
	// 可以看到无论是map还是struct都可以转成json数据，但是反之则不一定
	Struct2Map()
	fmt.Println(strings.Repeat("*", 100))
	Map2Struct() // 注意这里map的id字段是字符串
}

type MarsRover2 struct {
	Name       string `json:"name"`
	LaunchDate string `json:"launchDate"`
	Status     string `json:"status"`
}

func DemoJson() {
	// struct 写法
	var dataA []MarsRover2
	dataA = append(dataA, MarsRover2{Name: "Perseverance", LaunchDate: "2020-07-30", Status: "Active"})
	dataA = append(dataA, MarsRover2{Name: "Curiosity", LaunchDate: "2011-11-26", Status: "Active"})
	dataA = append(dataA, MarsRover2{Name: "Spirit", LaunchDate: "2003-06-10", Status: "Inactive"})
	// map 写法
	var dataB []map[string]string
	dataB1 := map[string]string{"name": "Perseverance", "launchDate": "2020-07-30", "status": "Active"}
	dataB2 := map[string]string{"name": "Curiosity", "launchDate": "2011-11-26", "status": "Active"}
	dataB3 := map[string]string{"name": "Spirit", "launchDate": "2003-06-10", "status": "Inactive"}
	dataB = append(dataB, dataB1, dataB2, dataB3)
	data := []MarsRover2{
		{Name: "Perseverance", LaunchDate: "2020-07-30", Status: "Active"},
		{Name: "Curiosity", LaunchDate: "2011-11-26", Status: "Active"},
		{Name: "Spirit", LaunchDate: "2003-06-10", Status: "Inactive"},
	}
	filename := "rovers.json"
	jsonData, _ := json.MarshalIndent(data, "", "    ") //格式化，空行
	_ = os.WriteFile(filename, jsonData, 0644)
	// 读取并追加json数据
	jsonFile, _ := os.Open(filename)
	defer jsonFile.Close()
	byteDate, _ := io.ReadAll(jsonFile)
	dSlice := new([]MarsRover2)
	_ = json.Unmarshal(byteDate, dSlice)
	for _, value := range *dSlice {
		fmt.Println(reflect.TypeOf(value), value)
	}
	*dSlice = append(*dSlice, MarsRover2{Name: "Natalia", LaunchDate: "2002-01-01", Status: "Active"})
	jsonData, _ = json.MarshalIndent(dSlice, "", "    ")
	_ = os.WriteFile(filename, jsonData, 0644)
	// 以追加模式和写权限打开文件(这会破坏Json结构)
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	// 追加字符串 "hello world" 到文件末尾
	_, _ = file.WriteString("\nhello world")
}
