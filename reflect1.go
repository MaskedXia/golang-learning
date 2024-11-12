package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectFun(arg interface{}) { //反射
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type resume struct {
	Name string `json:"name" doc:"My name"`
	Age  int    `json:"age" doc:"My age"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		doc := field.Tag.Get("doc")
		fmt.Printf("field: %s, json tag: %s, doc: %s\n", field.Name, tag, doc)
	}
}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year`
	Price  float64  `json:"price"`
	Actors []string `json:"act:"`
}

func main() {

	a := "hello"

	var allType interface{}
	allType = a
	fmt.Printf("type of allType: %T, value: %v\n", allType, allType)

	var num float32 = 1.23456
	reflectFun(num)

	var r resume
	findTag(r)

	fmt.Println("================================")

	// json marshal
	movie := Movie{"喜剧之王", 2000, 20, []string{"周星驰", "张柏芝"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	fmt.Println("jsonStr:", string(jsonStr))

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Printf("myMovie: %+v\n", myMovie)
}
