package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/borud/hav/pkg/auction"
)

const hetznerAuction = "https://www.hetzner.com/a_hz_serverboerse/live_data.json"

func main() {
	auction, err := getData()
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	for _, server := range auction.Server {
		fmt.Printf("name='%s' cpu='%s' ram='%d' price='%s' text='%s'\n",
			server.Name, server.CPU, server.RAM, server.Price, server.Freetext)
	}
}

func getData() (auction.Auction, error) {
	resp, err := http.Get(fmt.Sprintf("%s?%d", hetznerAuction, time.Now().UnixNano()/int64(time.Millisecond)))
	if err != nil {
		return auction.Auction{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var auction auction.Auction
	return auction, json.Unmarshal(body, &auction)
}
