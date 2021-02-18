package aboutPanicAndRecover

import (
	"errors"
	"fmt"
)

func trouble(i int){
	defer func() {
		if err:=recover();err!=nil{
			//恢复错误
			fmt.Printf("有错误产生%v",err)
		}
	}()
	if i!=1{
		panic(errors.New("i != 1\n"))
	}
}
