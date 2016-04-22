package main

import (
	"log"
	"os"

	sf "github.com/wwkeyboard/stockfighter"
)

func main() {
	conf, err := sf.ReadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	stockName := os.Args[1]

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}

	bid, ask, err := venue.GetSpread(stockName)
}
