package util

import (
	"fmt"
	"io"
	"os"
)

type Filehelper struct {
}

func (o *Filehelper) ReadTxt() string {
	var filename = "./grpcServer_out.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_RDWR, 0666) //打开文件
	} else {
		f, err1 = os.Create(filename) //创建文件
	}
	check(err1)

	fileinfo, _ := f.Stat()
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	bytesread, err := f.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(rune(bytesread))
}

func (o *Filehelper) WriteToTxt(str string, filename string) {
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_RDWR, 0666) //打开文件
	} else {
		f, err1 = os.Create(filename) //创建文件
	}
	check(err1)

	io.WriteString(f, str) //写入文件(字符串)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
