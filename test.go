package main

import (
	"fmt"
	"unsafe"
	"strings"
)

func getVal(num1 int,num12 int) (int){
	var sum = num1 + num12
	return sum
}


func main(){
	var sum = getVal(1000000,30)
	fmt.Println(sum)
	fmt.Println("hello world")
	fmt.Printf("num sum = %T size = %d", sum,unsafe.Sizeof(sum))
	a,b := swap("hello","world")
	fmt.Println(a,b)

	d := strings.Fields("helloworld")
	fmt.Println(d)

	arr := make(map[string]int) 

	for lenth,value := range d{
		_,ok := arr[value]
		if ok {
			fmt.Println(ok)
		}
		fmt.Println(ok)
		fmt.Println(value,lenth)
	}
}


func swap(x, y string)(string,string){
	return y,x
}
