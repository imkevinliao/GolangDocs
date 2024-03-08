package code

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
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
	var data map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Println("convert json to map failed!")
	}
	fmt.Println(data)
}

func DemoConvert() {
	Struct2Map()
	Map2Struct() // 注意这里map的id字段是字符串
}
