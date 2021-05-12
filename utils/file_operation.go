package utils

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {

	//打开文件
	file, err:=os.Open(filename)
	defer file.Close()

	//不存在则创建
	if err != nil && os.IsNotExist(err){
		file, _ = os.Create(filename)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//读取文件大小并创建缓冲区
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	//读取文件到缓冲区
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(buffer), nil
}

//向文件尾部追加字符串
func AppendToFile(file, str string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	f.WriteString(str)
}

//覆盖写入
func WriteToFile(file, str string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	f.WriteString(str)
}

