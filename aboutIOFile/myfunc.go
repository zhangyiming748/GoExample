package aboutIOFile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//创建一个新文件 golang.txt，并在其中写入 5 句“http://c.biancheng.net/golang/”
func Create(filePath string) {
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	//filePath := "e:/code/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("http://c.biancheng.net/golang/ \n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
//打开一个存在的文件，在原来的内容追加内容“C语言中文网”
func Appending(filePath string) {
	//filePath := "e:/code/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("C语言中文网 \r\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

//打开一个存在的文件，将原来的内容读出来，显示在终端，并且追加 5 句“Hello，C语言中文网”。
func ReadAndAppend(filePath string) {
	//filePath := "e:/code/golang.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("Hello，C语言中文网。 \r\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
//将一个文件的内容复制到另外一个文件（注：这两个文件都已存在）

func F2F(file1Path,file2Path string) {
	//file1Path := "e:/code/golang.txt"
	//file2Path := "e:/code/other.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
	}
}