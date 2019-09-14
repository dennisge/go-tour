package main

import (
	"fmt"

	"github.com/dennisge/go-tour/hello"
)

func main() {
	v := "Hello, 你好world"
	reverse := hello.Reverse(v)
	fmt.Println(reverse)

	const runes = "你好吗？"
	for index, runeValue := range runes {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

}
