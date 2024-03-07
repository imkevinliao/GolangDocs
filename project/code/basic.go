package code

import (
	"fmt"
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
