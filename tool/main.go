package main

import (
	"flag"
	"fmt"
	"github.com/hnit-acm/go-common/basic"
	"github.com/hnit-acm/go-common/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	flag.Parse()
	if !flag.Parsed() {
		return
	}
	args := flag.Args()
	fmt.Println(args)
	argsString := utils.ArrayStringToString(args, " ")
	exp, err := regexp.Compile(`^new \S+$`)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exp.MatchString(argsString) {
		fmt.Println("new service ", args[1])
		newService(basic.String(args[1]))
		return
	}

}

func newService(name basic.String) bool {
	fileList, _ := ioutil.ReadDir("./")
	for _, fileInfo := range fileList {
		// 如果存在模板文件
		if fileInfo.IsDir() && fileInfo.Name() == "template" {
			_, err := os.Open(name.GetNative())
			if err == nil {
				fmt.Println("服务已存在")
				return false
			}
			copyDir(fileInfo.Name(), name.GetNative())
			return true
		}
	}
	fmt.Println("不存在模板文件")
	return false
}

func copyDir(src, dest string) {
	fileList, _ := ioutil.ReadDir(src)
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileList {
		// 如果是目录
		if info.IsDir() {
			// 新建目录
			err := os.MkdirAll(filepath.Join(dest, info.Name()), os.ModePerm)
			if err != nil {
				return
			}
			copyDir(filepath.Join(src, info.Name()), filepath.Join(dest, info.Name()))
			continue
		}
		// 文件
		data, _ := ioutil.ReadFile(filepath.Join(src, info.Name()))
		err := ioutil.WriteFile(filepath.Join(dest, info.Name()), data, os.ModePerm)
		if err != nil {
			return
		}
		continue
	}
}
