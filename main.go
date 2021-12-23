package main

import "fmt"


func main() {
	fmt.Println("this is a stand program")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
