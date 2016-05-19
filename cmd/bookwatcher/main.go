package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"

	sf "github.com/wwkeyboard/stockfighter"
)

func socketURLs(c *sf.Configuration) (string, string) {
	return "http://api.stockfighter.io",
		fmt.Sprintf("wss://api.stockfighter.io/ob/api/ws/%s/venues/%s/executions/stocks/%s", c.VenueAccount, c.VenueName, c.Stock)
}

func main() {
	conFile := os.Args[1]
	conf, err := sf.ReadConfig(conFile)
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	origin, url := socketURLs(conf)
	fmt.Println("Dialing %s", url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	fmt.Println("Connected, Listening")

	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Println("have read")
	fmt.Printf("Received: %s.\n", msg[:n])
}
