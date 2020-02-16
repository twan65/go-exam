package main

func main() {
	var k int = 10
	var p = &k  // Kのアドレスを指す。
	println(*p) // pが指しているアドレスの実際値。
}
