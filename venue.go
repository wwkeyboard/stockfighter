package stockfighter

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Venue represents the stock trading venue
type Venue struct {
	Downloader Downloader
	Name       string
	Account    string
}

type heartbeatResponse struct {
	Ok    bool   `json:"ok"`
	Venue string `json:"venue"`
}

// MakeVenueForExchange takes the name of an exchange and returns a
// Venue configured for the production api
func MakeVenueForExchange(conf *Configuration) (*Venue, error) {
	downloader := HTTPDownloader{
		BaseURL: "https://api.stockfighter.io",
		Token:   conf.Token,
	}

	venue := Venue{
		Downloader: &downloader,
		Account:    conf.VenueAccount,
		Name:       conf.VenueName,
	}

	up, err := venue.IsUP()
	if err != nil {
		return nil, err
	}
	if !up {
		msg := fmt.Sprintf("The venue %s is down", conf.VenueName)
		return nil, errors.New(msg)
	}

	return &venue, nil
}

func (v Venue) heartbeatURL() string {
	return fmt.Sprintf("/ob/api/venues/%s/heartbeat", v.Name)
}

func (v Venue) orderURL(stock string) string {
	return fmt.Sprintf("/ob/api/venues/%s/stocks/%s/orders", v.Name, stock)
}

func (v Venue) orderbookURL(stock string) string {
	return fmt.Sprintf("/ob/api/venues/%s/stocks/%s", v.Name, stock)
}

func (v Venue) quoteURL(stock string) string {
	return fmt.Sprintf("/ob/api/venues/%s/stocks/%s/quote", v.Name, stock)
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

// Order is an outgoing order
type Order struct {
	Account   string `json:"account"`
	VenueName string `json:"venue"`
	Stock     string `json:"stock"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
	Direction string `json:"direction"`
	OrderType string `json:"orderType"`
}

// BuyLimit places a limit order
func (v Venue) BuyLimit(name string, price int, quantity int) (bool, error) {
	o := Order{
		Account:   v.Account,
		VenueName: v.Name,
		Stock:     name,
		Price:     price,
		Qty:       quantity,
		Direction: "buy",
		OrderType: "limit",
	}

	payload, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := v.Downloader.PostJSON(v.orderURL(name), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got: %s", resp)

	//	res := heartbeatResponse{}
	//	err = json.Unmarshal(resp, &res)
	//	if err != nil {
	//		return false, err
	//	}

	return true, nil

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

func (v Venue) GetQuoteAsk(stock string) (int, error) {
	resp, err := v.Downloader.GetJSON(v.quoteURL(stock))
	if err != nil {
		log.Fatal(err)
	}

	quote := make(map[string]interface{})
	err = json.Unmarshal(resp, &quote)
	if err != nil {
		return -1, err
	}
	ask, err := strconv.Atoi(quote["ask"].(string))
	if err != nil {
		return -1, err
	}

	return ask, nil
}
