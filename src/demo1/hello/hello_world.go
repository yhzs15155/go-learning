package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	if len(os.Args) > 1{
		fmt.Println("Hello world", os.Args[1])
	}

	var a = 1

	fmt.Println(a)

	// 循环
	for i := 0; i <= 10; i ++ {

	}

	os.Exit(-1)
}
