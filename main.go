package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string

	root := "D:/中医书籍/书籍"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Mode()&os.ModeDir > 0 {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})
	if err != nil {
		panic(err)
	}

	//file := os.OpenFile("readme.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)

	//for _, file := range files {
	//	fmt.Println(file)
	//}

	os.WriteFile("README.md", []byte("````\n"+
		strings.Join(files, "\r\n")+
		"\n````\n共"+fmt.Sprintf("%d", len(files))+"本\r\n### 加微信获取\n![alt 属性文本](w.png)"), 0777)
}
