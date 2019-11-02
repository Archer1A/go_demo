package main

import "fmt"

func main(){
	// 与java 不同 类型在变量名后
	var a string = "Runoob"
	fmt.Println(a)
	// 可以一次声明多个变量
	var b,c int = 1,2
	fmt.Println(b,c)

	var hello = "hello"
	//未初始化输出未空
	var world string
	fmt.Println(hello,world)

}