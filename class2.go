package main //程序的包名

import "fmt"

// 结构体
type Person struct {
	// 大写public 小写private
	name string
	age  int
}

func changePerson(p1 *Person) {
	p1.age = 20
    p1.name = "Jerry"
}

func (this *Person) GetName() string {
	return this.name
}

func (this *Person) SetName(name string) {
    this.name = name
}

func (sp *Person) test() {
    fmt.Println("Person test....") 
}

// 继承
type SuperPerson struct {
	Person
	level int
}

func (sp *SuperPerson) test() {
    fmt.Println("SuperPerson test....") 
}


type Animal interface{
	Sleep()
	GetColor() string
}

type Cat struct {
    color string
}

func (cat *Cat) Sleep(){
	fmt.Println("cat is sleep ...")
}

func (cat *Cat) GetColor() string{
	return cat.color
}

type Dog struct{
	color string
}

func (dog *Dog) Sleep(){
    fmt.Println("dog is sleep...")
}

func (dog *Dog) GetColor() string{
    return dog.color
}



func main() {
	var p1 Person
	p1.age = 12
	p1.name = "Tom"
	fmt.Println("p1 = ", p1) 
	changePerson(&p1)
	fmt.Println("p1 = ", p1) 

	p2 := Person{name: "Lily", age: 15}
	p2.SetName("Lisy")
	fmt.Println("p2 = ", p2.GetName())

	s1 := SuperPerson{Person{name: "Lily", age: 16}, 2}
	s1.test()
	fmt.Println("s1 = ", s1)  // 结构体嵌套，s1.Person.name = Lily, s1.age = 16, s1.level = 2

	var animal Animal
	animal = &Cat{"blank"}  //多态
	animal.Sleep()
	animal = &Dog{"yellow"}
	animal.Sleep()
}
