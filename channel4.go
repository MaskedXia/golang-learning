package main

import (
	"fmt"
	"time"
)

// 监控多路
func fibonacci(c, c2 chan int){
	x, y := 1, 1
	for {
		select {
        case c <- x:
            x, y = y, x+y
        case <-c2:
            close(c)
			fmt.Println("quit")
            return
        }
	}
}

func main() {
	c := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子 goroutine正在运行 ", <-c)
		}
		c2 <- 0
	}()

	time.Sleep(time.Second)

	fibonacci(c, c2)

	fmt.Println("主 goroutine结束...")
}