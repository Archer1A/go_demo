package main
import "fmt"

type Vertex struct{ // 结构体
	X int
	Y int
}

func main(){ // 结构体字段通过点号来访问
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := &v  // 结构体指针

	p.Y = 1e9 
	fmt.Println(v)

	k := v  // 结构体指针

	k.Y = 9 
	fmt.Println(v)
	fmt.Println(k)
}