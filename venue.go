package venue

import (
	"encoding/json"
	"fmt"
	"log"
)

// Venue represents the stock trading venue
type Venue struct {
	Downloader Downloader
	Name       string
}

type heartbeatResponse struct {
	Ok    bool   `json:"ok"`
	Venue string `json:"venue"`
}

func (v Venue) heartbeatURL() string {
	return fmt.Sprintf("/ob/api/venues/%s/heartbeat", v.Name)
}

func (v Venue) orderbookURL(stock string) string {
	return fmt.Sprintf("/ob/api/venues/%s/stocks/%s", v.Name, stock)
}

// IsUP returns true if the venue is up
func (v Venue) IsUP() (bool, error) {
	resp, err := v.Downloader.GetJSON(v.heartbeatURL())
	if err != nil {
		log.Fatal(err)
	}

	res := heartbeatResponse{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return false, err
	}

	return res.Ok, nil
}

// Bid represents a single bid on the orderbook
type Bid struct {
	Price int  `json:"price"`
	Qty   int  `json:"qty"`
	IsBuy bool `json:"isBuy"`
}

// OrderBook represents the orderbook for a specific stock
type OrderBook struct {
	OK     bool   `json:"ok"`
	Venue  string `json:"venue"`
	Symbol string `json:"symbol"`
	Bids   []Bid  `json:"bids"`
}

// GetBids queries the whole orderbook and returns the bids
func (v Venue) GetBids(stock string) ([]Bid, error) {
	resp, err := v.Downloader.GetJSON(v.orderbookURL(stock))
	if err != nil {
		log.Fatal(err)
	}

	book := OrderBook{}
	err = json.Unmarshal(resp, &book)
	if err != nil {
		return nil, err
	}

	return book.Bids, nil
}
