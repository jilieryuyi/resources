package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string

	root := "D:/中医书籍/中医书籍"
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

	os.WriteFile("README.md", []byte(
		"## 中医资料大全，持续更新"+
			"\n````\n"+
			strings.Join(files, "\r\n")+
			"\n````\n"+
			"共"+fmt.Sprintf("%d", len(files))+"本, 持续更新\r\n"+
			"### 加微信获取\n![alt 属性文本](w.png)"),
		0777)
}
