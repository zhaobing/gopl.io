package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	str = "hello世界"
	fmt.Printf("字符串:%s\nlen(str):%d,字节数量: %d, 字符数量: %d\n", str, len(str),
		len([]byte(str)), utf8.RuneCount([]byte(str)))

	i := len(str)
	i2 := int64(i)
	fmt.Println("i2", i2)
}
