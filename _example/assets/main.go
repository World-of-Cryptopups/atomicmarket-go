package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicmarket-go"
)

func main() {
	a := atomicmarket.New()

	get, err := a.GetAssets(&atomicmarket.GetAssetsQuery{
		Limit: 1,
	})

	fmt.Println(get)
	if err != nil {
		fmt.Println("Err")
		log.Fatal(err)
	}
	id := get.Data[0].AssetID

	fmt.Println(get.Data[0].Name, " - ", id)

	stats, err := a.GetAssetIDStats(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stats.Data.TemplateMint)

	logs, err := a.GetAssetIDLogs(id, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(logs.Data[0].Name)

	sales, err := a.GetAssetIDSales(id, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sales)

	if len(sales.Data) > 0 {
		fmt.Println(sales.Data[0].SaleID)
	}
}
