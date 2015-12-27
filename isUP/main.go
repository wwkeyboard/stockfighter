package main

import (
	"fmt"
	"log"
	"os"

	sf "github.com/wwkeyboard/stockfighter"
)

func main() {
	conf, err := sf.ReadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	fmt.Printf("testing venue %s\n", conf.VenueName)

	venue, err := sf.MakeVenueForExchange(conf)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", conf.VenueName, err)
	}
	venue.IsUP()
}
