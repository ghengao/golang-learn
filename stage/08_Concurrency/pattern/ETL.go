package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

//Extract, Transform, Load

func main() {
	start := time.Now()
	orders := extract()
	orders = transform(orders)
	load(orders)
	fmt.Println(time.Since(start))
}

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerId int
	PartNumber string
	Qty        int
	UnitCost   float64
	UnitPrice  float64
}

func extract() []*Order {
	result := []*Order{}
	ordFh, _ := os.Open("./orders.txt")
	defer ordFh.Close()
	r := csv.NewReader(ordFh)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerId, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Qty, _ = strconv.Atoi(record[2])
		append(result, order)
	}
	return result
}

func transform(orders []*Order) []*Order {
	prodFh, _ := os.Open("./productList.txt")
	defer prodFh.Close()
	r := csv.NewReader(prodFh)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}
	for idx := range orders {
		time.Sleep(3 * time.Millisecond) // Simulate Web request delay
		o := orders[idx]
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
	}
	return orders

}

func load(orders []*Order) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	for _, o := range orders {
		time.Sleep(1 * time.Millisecond) // Simulate the lag to write to the db
		fmt.Fprintf(f, "%20s %15d %12.2f %12;2f %15.2f %15.2f\n", o.PartNumber, o.Qty, o.UnitCost, o.UnitPrice, o.UnitCost*float64(o.Qty), o.UnitPrice*float64(o.Qty))
	}
}
