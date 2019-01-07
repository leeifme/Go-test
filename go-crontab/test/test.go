package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

//var ch = make(chan int)
var mutex sync.Mutex // 创建一个互斥锁（互斥量）

func printer(s string) {
	mutex.Lock()         // 访问共享数据前 加锁。
	defer mutex.Unlock() // 访问共享数据结束，解锁。
	for _, char := range s {
		fmt.Printf("%c", char) // stdout
		time.Sleep(time.Second * 1)
	}
	fmt.Println("")
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	// go person1()
	// go person2()

	// for {
	// }
	person := Person{
		Name:   "Wang Wu",
		Age:    30,
		Weight: 150.07,
	}

	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes))

	var personFromJSON interface{}
	json.Unmarshal(jsonBytes, &personFromJSON)
	fmt.Println(personFromJSON)
	fmt.Println(reflect.TypeOf(personFromJSON))

	r := personFromJSON.(map[string]interface{})

	fmt.Println(reflect.TypeOf(r["Age"]).Name())    // float64
	fmt.Println(reflect.TypeOf(r["Weight"]).Name()) // float64
}
