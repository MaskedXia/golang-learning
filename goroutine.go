package main

import (
    "fmt"
	"time"
)

// 从 goroutine
func newTask(){
	i := 0
	for {
		i++
        fmt.Printf("Task %d\n", i)
        time.Sleep(time.Second)
	}
}


// 主 goroutine
func main() {
	fmt.Println("Starting")

	go newTask()
	i := 0

	for {
        i++
        fmt.Printf("Main %d\n", i)
        time.Sleep(time.Second)
    }
}