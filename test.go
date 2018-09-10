package main

import "fmt"

func main() {
	type account struct {
		value int    // 토큰 량
		kind  string // user, store 구분
	}
	var BalanceOf = make(map[string]account)
	BalanceOf["test"] = account{value: 20, kind: "store"}
	fmt.Println(BalanceOf["test"].value)
	fmt.Println(BalanceOf["test"].kind)
}
