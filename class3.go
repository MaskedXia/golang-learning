package main

import "fmt"

//interface{}万能数据类型
func fun1(arg interface{}){
	fmt.Println("fun1 called")
	fmt.Println(arg)

	//断言机制
	switch v := arg.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    default:
        fmt.Printf("unknown type\n")
    }
}

type Book struct {
	Title string
}


func main() {
	book := Book{"Golang"}
	fun1(book)
}