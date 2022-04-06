package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicmarket-go"
)

func main() {
	a := atomicmarket.New()

	get, err := a.GetTransfers(&atomicmarket.GetTransfersQuery{
		Limit: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(get.Data[0].TransferID)
}
