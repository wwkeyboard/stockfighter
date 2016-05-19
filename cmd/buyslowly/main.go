package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	sf "github.com/wwkeyboard/stockfighter"
)

func main() {
	conFile := os.Args[1]
	conf, err := sf.ReadConfig(conFile)
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}

	stockName := os.Args[2]
	qty := os.Args[3]

	stockQuantity, err := strconv.Atoi(qty)
	if err != nil {
		log.Fatalf("Could not parse quantity %s, %s", qty, err)
	}

	loop(venue, stockName, stockQuantity)
}

func loop(venue *sf.Venue, stockName string, wantQty int) {
	needQty := wantQty
	boughtQty := 0
	// while boughtQty < wantQty
	//for boughtQty < wantQty {
	price, err := venue.GetQuoteAsk(stockName)

	// place order
	currentOrder, err := venue.BuyLimit(stockName, price+1, needQty)
	if err != nil {
		log.Fatalf("error buying stock %s", err)
	}

	// wait
	time.Sleep(1000)
	fmt.Println("looking for %s", currentOrder)
	// close unfilled order
	// boughtQty = boughtQty + filledQty
	//}
	fmt.Println("new boughtQty %s", boughtQty)
}
