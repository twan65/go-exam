package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "Lee" // p가 포인터라도 . 을 사용한다
	test(p)
	fmt.Println(p)
}

func test(param person) {
	param.age = 20
}
