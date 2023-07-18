package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files, other []string

	filepath.Walk("D:/中医书籍/中医书籍", func(path string, info os.FileInfo, err error) error {
		if info.Mode()&os.ModeDir > 0 {
			return nil
		}

		size := info.Size()
		size = size / 1024
		sizeText := fmt.Sprintf("%dk", size)
		if size > 1024 {
			size = size / 1024
			sizeText = fmt.Sprintf("%dM", size)
		}

		files = append(files, info.Name()+" "+sizeText)
		return nil
	})

	filepath.Walk("D:/中医书籍/其他电子书", func(path string, info os.FileInfo, err error) error {
		if info.Mode()&os.ModeDir > 0 {
			return nil
		}

		size := info.Size()
		size = size / 1024
		sizeText := fmt.Sprintf("%dk", size)
		if size > 1024 {
			size = size / 1024
			sizeText = fmt.Sprintf("%dM", size)
		}

		other = append(other, info.Name()+" "+sizeText)
		return nil
	})

	os.WriteFile("README.md", []byte(
		"## 中医资料大全，持续更新"+
			"\n````\n"+
			strings.Join(files, "\r\n")+
			"\n````\n"+
			"共"+fmt.Sprintf("%d", len(files))+"本, 持续更新\r\n"+
			"### 加微信获取\n![alt 属性文本](w.png)\n## 其他电子书\n````\n"+
			strings.Join(other, "\n")+
			"\n````\n"),
		0777)
}
