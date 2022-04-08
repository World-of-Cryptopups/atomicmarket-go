# atomicmarket-go

An Api wrapper for the AtomicMarket API.

## Note
All functions are not tested or guaranteed to work correctly yet.

## Install 
```
go get -u github.com/World-of-Cryptopups/atomicmarket-go
```


## Usage
```go
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

}
```


##
**&copy; 2022 | World of Cryptopups**