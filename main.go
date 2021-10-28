package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("this is a stand program")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	p1 := person{
		name: "zhangyiming",
		age:  26,
	}
	p2 := person{
		name: "zhangyiming",
		age:  26,
	}
	if p1 == p2 {
		fmt.Println("same people")
	} else {
		panic("结构体不能比较")
	}
}
