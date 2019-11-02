package main

import(
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	result :=make( map[string]int)
	for _,value := range strings.Fields(s){
		result[value] += 1
		/*_,ok := result[value]
		if ok {
			result[value] = result[value] + 1
		}else{
			result[value] = 1
		}*/
	}
	return result
}

func main(){
	fmt.Println(WordCount("hello world"))
}

