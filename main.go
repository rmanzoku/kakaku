package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var bitbankAPI = "https://public.bitbank.cc/"

type Ticker struct {
	Success int64 `json:"success"`
	Data    Data  `json:"data"`
}

type Data struct {
	Sell      string `json:"sell"`
	Buy       string `json:"buy"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Last      string `json:"last"`
	Vol       string `json:"vol"`
	Timestamp int64  `json:"timestamp"`
}

func FetchTicker(pair string) (*Ticker, error) {
	response, err := http.Get(bitbankAPI + pair + "/ticker")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var ret = new(Ticker)
	if err := json.Unmarshal(bytes, ret); err != nil {
		return nil, err
	}

	return ret, nil

}

func main() {

	btcJpyTicker, err := FetchTicker("btc_jpy")
	ethBtcTicker, err := FetchTicker("eth_btc")

	if err != nil {
		log.Fatal(err)
	}

	btcJpy, err := strconv.ParseFloat(btcJpyTicker.Data.Buy, 64)
	if err != nil {
		log.Fatal(err)
	}

	ethBtc, err := strconv.ParseFloat(ethBtcTicker.Data.Buy, 64)
	if err != nil {
		log.Fatal(err)
	}

	k := "pricing.eth.jpy"
	v := ethBtc * btcJpy

	fmt.Printf("%s\t%f\t%d\n", k, v, time.Now().Unix())
}
