package main

import (
	"fmt"
	"time"

	socket "github.com/canhlinh/tradingview-scraper/v2"
)

func main() {
	tradingviewsocket, err := socket.Connect(
		func(symbol string, data *socket.QuoteData) {
			if data.Price != nil {

				fmt.Printf("%v %v \n", *data.Price, *data.Timestamp)
			}
		},
		func(err error, context string) {
			fmt.Printf("%#v", "error -> "+err.Error())
			fmt.Printf("%#v", "context -> "+context)
		},
	)
	if err != nil {
		panic("Error while initializing the trading view socket -> " + err.Error())
	}
	tradingviewsocket.AddSymbol("BINANCE:BTCUSDT")
	for {
		<-time.After(time.Second)
	}
}
