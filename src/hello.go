package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	n := 1
	for n < 100 {
		n *= 2
		//if n > 90 {
		//   break
		//}
	}
	println(n)

	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}
