package main

import (
	"fmt"
	"os"
	"readethnode/enode"
	"time"
)

func main() {
	param1:=os.Args[1]
	db, err := enode.OpenDB(param1)
	if err != nil {
		fmt.Println(err)
	}
	seeds := db.QuerySeeds(500, 500*24*time.Hour)
	for index, value := range seeds{
		fmt.Printf("Index: %d IP: %d ID: %s Seq: %d UDP: %d\n", index, value.IP(),value.ID().String(),value.Seq(),value.UDP())
	}

}
