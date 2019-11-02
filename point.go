package main
import "fmt"
func main(){
	i,j := 42, 2701
	p := &i // & 操作符会生成一个指向其操作数的指针
	fmt.Println(*p) // 通过指针p读取i 
	*p = 21
	fmt.Println(i) 

	p = &j
	*p = *p / 37
	fmt.Println(j)
}