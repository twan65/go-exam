package main

func main() {
	const a int = 10
	const b = 20
	const s = "Hi"
	const f = 11.

	const (
		Apple = iota
		Grape
		Orange
	)

	println(a)
	println(b)
	println(s)
	println(f)
	println(Apple)
	println(Grape)
	println(Orange)
}
