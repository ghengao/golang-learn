package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ordFh, err := os.Create("orders.txt")
	check(err)
	defer ordFh.Close()
	productFh, err := os.Create("productList.txt")
	check(err)
	defer ordFh.Close()
	for i := 0; i < 1000; i++ {
		customerId := 1000 + rand.Intn(8999)
		productId := 10000000 + rand.Int31n(89999999)
		Qty := 1 + rand.Intn(99)
		ordPrice := rand.Float64() * 100
		salePrice := ordPrice * 0.8
		ordLine := fmt.Sprintf("%d,%d,%d\n", customerId, productId, Qty)
		ordFh.WriteString(ordLine)
		prodLine := fmt.Sprintf("%d,%.2f,%.2f\n", productId, salePrice, ordPrice)
		productFh.WriteString(prodLine)
	}
}
