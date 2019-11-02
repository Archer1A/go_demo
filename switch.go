package main
import (
	"fmt"
	"runtime"
	"time"
)

func main(){

	windows := "windows"
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os{
	case "darwin" :
		fmt.Printf( "OS X")
	case "linux":
		fmt.Printf("Linux")
	case windows:
		fmt.Printf(windows)
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday{
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	} 

}