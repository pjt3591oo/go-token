package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", n, n)
	}
	fmt.Println(n)
}
