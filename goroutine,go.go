package main

import (
	"fmt"
	"runtime"
)

func say(s string)  {
	for i := 0;i<5 ;i++  {
		//表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		runtime.Gosched()

		fmt.Println(s)

	}

}

func sum(a []int ,c chan int)  {
	total := 0
	for _,v := range a{
		total += v
	}
	c <- total // send total to c

}

func main()  {
	go say("world")
	say("hello")
	fmt.Println(runtime.NumCPU())
	a := []int {7,2,8,-9,4,0}
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x , y := <-c ,<-c
	fmt.Println(x,y,x+y)

}
