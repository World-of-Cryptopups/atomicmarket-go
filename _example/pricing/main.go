package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicmarket-go"
)

func main() {
	a := atomicmarket.New()

	get, err := a.GetPriceAssets(&atomicmarket.GetPriceAssetsQuery{
		Owner: "5g2vm.wam",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(get.Data[0].SuggestedMedian)
}
