package main

func main() {
	msg := "world"
	say1(msg)  // Pass By Value
	say2(&msg) // Pass By Reference
}

func say1(msg string) {
	println("Hello, " + msg)
}

func say2(msg *string) {
	*msg = "Anveloper"
	println(*msg)
}
