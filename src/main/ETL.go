package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//Extract, Transform, Load

func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)
	go extract(extractChannel)
	go transform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)
	<-doneChannel
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

func extract(ch chan *Order) {
	//result := []*Order{}
	ordFh, _ := os.Open("./orders.txt")
	defer ordFh.Close()
	r := csv.NewReader(ordFh)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerId, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Qty, _ = strconv.Atoi(record[2])
		//result = append(result, order)
		ch <- order
	}
	//fmt.Println("Closing extract")
	close(ch)
	//return result
}

func transform(extractChannel, transformChannel chan *Order) {
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
	wg := sync.WaitGroup{}
	for o := range extractChannel {
		wg.Add(1)
		go func(o *Order) {
			defer wg.Done()
			time.Sleep(3 * time.Millisecond) // Simulate Web request delay
			//o := orders[idx]
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
		}(o)
	}
	wg.Wait()
	//fmt.Println("Closing transport")
	close(transformChannel)

}

func load(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	wg := sync.WaitGroup{}
	for o := range transformChannel {
		wg.Add(1)
		go func(o *Order) {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond) // Simulate the lag to write to the db
			fmt.Fprintf(f, "%20s%15d%12.2f%12.2f%15.2f%15.2f\n",
				o.PartNumber,
				o.Qty,
				o.UnitCost,
				o.UnitPrice,
				o.UnitCost*float64(o.Qty),
				o.UnitPrice*float64(o.Qty))
		}(o)
	}
	wg.Wait()
	//fmt.Println("Closing load")
	doneChannel <- true
}
