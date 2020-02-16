package main

func main() {
	say("This", "is", "a", "book")
	say("Hi")

	total := sum1(1, 2, 3, 4, 5, 6, 7)
	println(total)

	sum, cnt := getSumAndCnt(5, 4, 12, 32, 12)
	println(sum, cnt)
}

func say(msg ...string) {
	for _, s := range msg { // "_"はIndexを使用しないという意味。
		println(s)
	}
}

func sum1(nums ...int) int {
	s := 0

	for _, num := range nums {
		s += num
	}
	return s
}

func getSumAndCnt(nums ...int) (int, int) {
	s := 0
	cnt := 0

	for _, num := range nums {
		s += num
		cnt++
	}

	return s, cnt
}
