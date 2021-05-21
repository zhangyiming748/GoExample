package main

import "log"

func init() {
	log.SetPrefix("from: ")
	log.SetFlags(log.Ldate|log.Llongfile|log.Ltime)
}
func main() {
	defer func() {
		if err:=recover();err!=nil{
			log.Println(err)
		}
	}()
	log.Println("信息")

	log.Panicln("panic 信息")

	log.Fatalln("fatal 信息")
}
