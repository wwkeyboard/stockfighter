package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"

	sf "github.com/wwkeyboard/stockfighter"
)

func socketURLs(c *sf.Configuration) (string, string) {
	return "http://api.stockfighter.io",
		fmt.Sprintf("wss://api.stockfighter.io/ob/api/ws/%s/venues/%s/tickertape/stocks/%s", c.VenueAccount, c.VenueName, c.Stock)
}

func main() {
	conFile := os.Args[1]
	conf, err := sf.ReadConfig(conFile)
	if err != nil {
		log.Fatalf("Could not parse config: %s", err)
	}

	origin, url := socketURLs(conf)
	fmt.Printf("Dialing %s\n", url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	fmt.Println("Connected, Listening")

	// loop and list
	r := bufio.NewReaderSize(ws, 1536)
	line, _, err := r.ReadLine()

	scanner := bufio.NewScanner(ws)

	fmt.Println("------ received %v", string(line))
	for scanner.Scan() {
		fmt.Println("====== found")
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("------ Done")

}
