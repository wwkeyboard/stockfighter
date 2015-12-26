package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/wwkeyboard/stockfighter"
)

func makeVenueForExchange(name string) (*venue.Venue, error) {
	downloader := venue.HTTPDownloader{
		BaseURL: "https://api.stockfighter.io",
	}

	venue := venue.Venue{
		Downloader: &downloader,
		Name:       name,
	}

	up, err := venue.IsUP()
	if err != nil {
		return nil, err
	}
	if !up {
		msg := fmt.Sprintf("The venue %s is down", name)
		return nil, errors.New(msg)
	}

	return &venue, nil
}

func main() {
	venueName := os.Args[1]
	fmt.Printf("testing venue %s\n", venueName)

	venue, err := makeVenueForExchange(venueName)
	if err != nil {
		log.Fatalf("The venue %s is down with error %s", venueName, err)
	}
	venue.IsUP()
}
