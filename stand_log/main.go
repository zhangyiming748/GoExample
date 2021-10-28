package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//func init() {
//	log.SetPrefix("from: ")
//	log.SetFlags(log.Ldate | log.Llongfile | log.Ltime)
//}
//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(err)
//		}
//	}()
//	log.Println("信息")
//	log.Panicln("panic 信息")
//	log.Fatalln("fatal 信息")
//}
var (
	Trace   *log.Logger
	Info    *log.Logger
	Debug   *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	DebugLog, err1 := os.OpenFile("debug.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 != nil {
		log.Println("Debug日志文件打开失败")
	}
	ErrorLog, err2 := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		log.Println("Error错误记录打开失败")
	}
	Trace = log.New(ioutil.Discard, "TRACE:", log.LstdFlags|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lshortfile)
	Debug = log.New(io.MultiWriter(DebugLog, os.Stdout), "DEBUG:", log.LstdFlags|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING:", log.LstdFlags|log.Lshortfile)
	Error = log.New(io.MultiWriter(ErrorLog, os.Stderr), "ERROR:", log.LstdFlags|log.Llongfile)
}
func main() {
	Trace.Println("你说什么风太大我听不见")
	Info.Println("我随便说你随便听")
	Debug.Println("你有问题")
	Warning.Println("FBI WARNING")
	Error.Println("我不听我不听都是你的错")
}
