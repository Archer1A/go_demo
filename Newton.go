package main

import (
	"fmt"
	"math"
)
/**
* x 被开方的数
* y 循环次数  次数越多越准确
**/
func newton(x float64,y int) float64{
	z := 1.0
	for i := 0;i<y ;i++{
		z -= (z*z-x) / (x*2)
	}
	return z


}

func main(){
	for i :=0 ; i<10 ;i++{
		fmt.Println(i,"次结果----->",newton(2,i))
	}
	fmt.Println("math 包的计算结果",math.Sqrt(2))
	
}