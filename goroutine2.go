package main

import (
	"fmt"
	//"runtime"
	"time"
)

func main() {
	fmt.Println("Starting")

	// go func(){ //匿名函数
	// 	defer fmt.Println("A.defer")
	// 	func() {
	// 		defer fmt.Println("B.defer")

	// 		// 退出goroutine
	// 		runtime.Goexit()
    //         fmt.Println("B")
    //     }()

	// 	fmt.Println("A")
	// }()

	go func(a int, b int) int{
		fmt.Println("Sum:", a+b)
		return a + b
    }(10, 20)

	for {
        time.Sleep(time.Second)
    }
}