package main

import (
	"encoding/json"
	"fmt"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go データ
	mem := Member{"Alex", 10, true}

	// JSON インコ－ディン
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	// JSON byte→文字列
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}
