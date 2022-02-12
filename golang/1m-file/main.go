package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var alphabetical []string
	for i := 97; i <= 122; i++ {
		alphabetical = append(alphabetical, fmt.Sprintf("%c", i))

	}
	file, err := os.OpenFile("1mfile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	alphabeticalStr := strings.Join(alphabetical, "")
	count := 0
	for count <= 40329 {
		_, err := file.WriteString(alphabeticalStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		count++
	}
	// 简单版
	// fileInfo, _ := os.Stat("1mfile")
	// fmt.Printf("the size of 1mfile is %d", fileInfo.Size())
	//复杂的方案，更适合用来遍历一个文件夹
	err = filepath.Walk("1mfile", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		filesize := info.Size()
		fmt.Printf("the size of 1mfile is %d", filesize)
		return nil
	})
}
