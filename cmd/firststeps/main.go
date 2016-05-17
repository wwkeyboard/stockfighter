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

	stockName := os.Args[2]

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}

	log.Printf("%s - %s", venue, stockName)

	stockPrice, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Need to specify a price to buy")
	}
	success, err := venue.BuyLimit(stockName, stockPrice, 100)
	if err != nil {
		log.Fatalf("Failed to buy %s", err)
	}

	if !success {
		log.Fatalf("failed to buy")
	}
}
