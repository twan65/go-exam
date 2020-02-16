package main

import "fmt"

func main() {
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 出力
	println(len(target), cap(target)) // 3, 6 出力
}
