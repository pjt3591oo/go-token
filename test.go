package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	// txdb, _ := leveldb.OpenFile("db/transaction", nil)
	// data, _ := txdb.Get([]byte("3506b7a8b0b309bf5a6a19407d1d89bb933e5054611421ec795a4c810913206d"), nil)
	// defer txdb.Close()
	// fmt.Println(string(data))

	accountdb, _ := leveldb.OpenFile("db/account", nil)
	data, _ := accountdb.Get([]byte("6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"), nil)
	defer accountdb.Close()
	fmt.Println(string(data))

}
