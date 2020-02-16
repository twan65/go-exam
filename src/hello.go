package main

//Rect - struct定義
type Rect struct {
	width, height int
}

//Rectのarea()メソッド
func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20} // struct生成
	area := rect.area()  // メソッドの呼び出し
	println(area)
}
