package main

import (
	"fmt"
	//"runtime"
	//"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束...")
		fmt.Println("goroutine正在运行...")
		c <- 1
	}()

	num := <-c //会阻塞

	fmt.Println("主 goroutine接收到:", num)
	fmt.Println("主 goroutine结束...")
}