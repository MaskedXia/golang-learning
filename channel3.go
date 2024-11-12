package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子 goroutine正在运行，发送的元素是 ", i, " len(c) = ", len(c), " cap(c) =", cap(c))
		}
		// 关闭channel
		close(c)
	}()

	time.Sleep(time.Second)

	// for {
	// 	// 接收数据，ok为true表示接收到数据，为false表示channel已关闭
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	}else{
	// 		break // 接收完所有数据并关闭 channel，循环结束
	// 	}
	// }

	// 接收数据，直到channel被关闭为止，和上面功能相同
	for data := range c {
		fmt.Println(data)
    }

	fmt.Println("主 goroutine结束...")
}