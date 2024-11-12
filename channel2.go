package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带缓冲 
	// channel满写数据会阻塞，channel空取数据会阻塞
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子 goroutine结束...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子 goroutine正在运行，发送的元素是 ", i, " len(c) = ", len(c), " cap(c) =", cap(c))
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("主 goroutine接收到:", num)
	}

	fmt.Println("主 goroutine结束...")
}