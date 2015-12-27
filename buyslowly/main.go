package main

import (
	"log"
	"os"
	"strconv"

	"github.com/wwkeyboard/stockfighter"
)

func main() {
	venueName := os.Args[1]
	venueAccount := os.Args[2]
	stockName := os.Args[3]
	price := os.Args[4]
	qty := os.Args[5]

	stockPrice, err := strconv.Atoi(price)
	if err != nil {
		log.Fatalf("Could not parse price %s, %s", price, err)
	}
	stockQuantity, err := strconv.Atoi(qty)
	if err != nil {
		log.Fatalf("Could not parse quantity %s, %s", qty, err)
	}

	venue, err := venue.MakeVenueForExchange(venueName, venueAccount)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", venueName, err)
	}

	success, err := venue.BuyLimit(stockName, stockPrice, stockQuantity)
	if err != nil {
		log.Fatalf("error buying stock %s", err)
	}
	log.Printf("success was %s", success)
}
