package main //程序的包名

import (
	"fmt"
	"time"
)

// 声明函数
func testFun(a string, b string) (bool, int) {
	fmt.Println("a = ", a, " b = ", b)
	c := true
	return c, 100
}

func changeValue(p *int) {
	*p = 10
}

func changeArray(arr []int) {
	arr[0] = 1000
}

func main() {
	fmt.Println("Hello Go!")
	time.Sleep(1 * time.Second)

	//变量
	var a int = 10
	var b string = "Hello"
	var c bool = true
	fmt.Printf("a: %d, b: %s, c: %t\n", a, b, c)

	// 省略类型
	var d = 100
	fmt.Printf("d: %d, type of c is %T\n", d, d)

	//省略var
	e := 200
	fmt.Printf("e: %d, type of e is %T\n", e, e)

	// 多重赋值
	f, g := 300, "World"
	fmt.Printf("f: %d, g: %s\n", f, g)

	// 常量
	const pi = 3.141592653589793
	fmt.Printf("pi: %f, type of pi is %T\n", pi, pi)

	// iota 常量生成器
	const (
		x = iota // 0
		y        // 1
		z        // 2
	)
	fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)

	ret1, ret2 := testFun("111", "222")
	fmt.Println("ret1:", ret1, "ret2:", ret2)

	//指针
	var ac int = 1
	changeValue(&ac)
	fmt.Println("ac after changeValue:", ac)

	defer fmt.Println("defer") // 最为最后执行，类似栈
	fmt.Println("defer test")

	var myArrays []int

	for i := 0; i < 10; i++ {
		myArrays = append(myArrays, i)
	}
	for index, value := range myArrays {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	fmt.Println("================================")
	myArrays2 := [10]int{1, 2, 3}
	for i := 0; i < len(myArrays2); i++ {
		fmt.Println(myArrays2[i])
	}

	fmt.Println("================================")
	myArrays3 := []int{10, 20, 30} //动态数组，引用传递
	for _, v := range myArrays3 {
		fmt.Println(v)
	}
	fmt.Println("===========after change===============")
	changeArray(myArrays3)
	fmt.Printf("len = %d, myArrays3 = %v\n", len(myArrays3), myArrays3)

	fmt.Println("================================")
	slice := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice), slice)
	slice = nil  //空切片
	fmt.Printf("len = %d, slice = %v\n", len(slice), slice)

	fmt.Println("================================")
	slice2 := make([]int, 3, 5) //容量是5
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(slice2), cap(slice2), slice2)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 2)
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(slice2), cap(slice2), slice2)

	fmt.Println("================================")
	slice3 := make([]int, 3, 5) 
	s3 := slice3[0:2]  //和python一样的切片
	fmt.Printf("len = %d, cap = %d, s3 = %v\n", len(s3), cap(s3), s3)
	s4 := []int{1, 2, 3}
	copy(s3, s4)  //拷贝
	fmt.Printf("len = %d, cap = %d, s3 = %v\n", len(s3), cap(s3), s3)
	fmt.Println("s3 = ", s3)

	fmt.Println("================================")
	map1 := make(map[string]int, 10) //
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println("map1 = ", map1)
	
	map2 := map[string]int{
		"a": 1,
        "b": 2,
	}
	fmt.Println("map2 = ", map2)

	for key, value := range map2{
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	delete(map2, "a")
	map2["b"] = 3
	fmt.Println("map2 after delete a and change value = ", map2)
}
