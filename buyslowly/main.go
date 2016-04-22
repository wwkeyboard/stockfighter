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

	stockName := os.Args[1]
	price := os.Args[2]
	qty := os.Args[3]

	stockPrice, err := strconv.Atoi(price)
	if err != nil {
		log.Fatalf("Could not parse price %s, %s", price, err)
	}
	stockQuantity, err := strconv.Atoi(qty)
	if err != nil {
		log.Fatalf("Could not parse quantity %s, %s", qty, err)
	}

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}

	success, err := venue.BuyLimit(stockName, stockPrice, stockQuantity)
	if err != nil {
		log.Fatalf("error buying stock %s", err)
	}
	log.Printf("success was %s", success)
}
