package main

import (
	"fmt"
	"os"
	"readethnode/enode"
)

func main() {
	param1:=os.Args[1]
	db, err := enode.OpenDB(param1)
	//db, err := enode.OpenDB("G:\\work\\区块链\\以太坊\\temp\\nodes")
	if err != nil {
		fmt.Println(err)
	}
	db.FindAll()
}
