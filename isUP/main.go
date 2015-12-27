package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wwkeyboard/stockfighter"
)

func main() {
	venueName := os.Args[1]
	fmt.Printf("testing venue %s\n", venueName)

	venue, err := venue.MakeVenueForExchange(venueName)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", venueName, err)
	}
	venue.IsUP()
}
