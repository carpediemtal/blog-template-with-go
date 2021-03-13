package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	title    string
	date     string
	tmplPath string
	postPath string
)

func init() {
	titlePtr := flag.String("title", "panic", "post name: LeetCode-???")
	tmplPtr := flag.String("tmpl", "C:\\Users\\67460\\Documents\\Golang\\leetcode-template\\resources\\template", "template file path")
	postPtr := flag.String("post", "C:\\Users\\67460\\Documents\\blog\\source\\_posts", "post dir path")
	flag.Parse()

	if *titlePtr == "panic" {
		panic("post title is required!")
	}
	title = *titlePtr
	tmplPath = *tmplPtr
	postPath = *postPtr
	date = time.Now().String()[:len("2021-03-13 10:08:40")]
}

func main() {
	filePath := fmt.Sprintf("%s/%s.md", postPath, title)
	if _, err := os.Stat(filePath); os.IsExist(err) {
		panic("file already exists")
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(getTmpl())
	if err != nil {
		panic(err)
	}
}

func getTmpl() string {
	tmpl, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(string(tmpl), title, date)
}
