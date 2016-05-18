package main

import (
	"log"
	"os"
	"strconv"

	sf "github.com/wwkeyboard/stockfighter"
)

func main() {
	conf, err := sf.ReadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}

	stockName := os.Args[1]
	qty := os.Args[2]

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
	for boughtQty < wantQty {
		price, err := venue.GetQuoteAsk(stockName)
		// place order
		success, err := venue.BuyLimit(stockName, price+1, needQty)
		if err != nil {
			log.Fatalf("error buying stock %s", err)
		}

		// wait
		// close unfilled order
		// boughtQty = boughtQty + filledQty
	}
}
