package main

import (
	"aboutGetVal/util"
	"fmt"
	"strings"
)

func main() {
	ret := util.GetVal("target", "cmd")
	fmt.Println(ret)
	s:=strings.Split(ret, " ")
	for i, v := range s {
		fmt.Printf("第%d段文字为%v\n", i, v)
	}
}
