package main

import "fmt"

type  Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human // 匿名字段 human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}
// human 对象实现say hi 接口
func (h *Human)SayHi()  {
	fmt.Printf("Hi , I am %s you can call me on %s \n",h.name,h.phone)

}
// human 实现sing 接口
func (h *Human)Sing(lyrics string)  {
	fmt.Printf("La la la la la la la ..." ,lyrics)

}
//hum 实现guzzle
func (h *Human)Guzzle(beerStein string)  {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e *Employee) SayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行

}
//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan  += amount // (again and again and...)
}
//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i Men

	i = &mike
	fmt.Println("this is Mike ,a Student:")
	i.SayHi()
	i.Sing("November rain")
	//i也能存储Employee
	i = &tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = &paul, &sam, &mike

	for _, value := range x{
		value.SayHi()
	}

}