package main

func main() {
	var k int = 1
	if k == 1 { // {が同じライン
		println("One")
	}

	k = 3
	if k == 1 {
		println("One")
	} else if k == 2 { // {が同じライン
		println("Two")
	} else { // {が同じライン
		println("Other")
	}
}
