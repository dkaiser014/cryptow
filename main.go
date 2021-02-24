package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"cryptow/pkg/loadenv"
)

// Crypto represents the data structure of the
// response returned from the coinmarketcap api
type Crypto struct {
	Data []struct {
		CirculatingSupply int64  `json:"circulating_supply"`
		CmcRank           int64  `json:"cmc_rank"`
		DateAdded         string `json:"date_added"`
		ID                int64  `json:"id"`
		LastUpdated       string `json:"last_updated"`
		MaxSupply         int64  `json:"max_supply"`
		Name              string `json:"name"`
		NumMarketPairs    int64  `json:"num_market_pairs"`
		Platform          struct {
			ID           int64  `json:"id"`
			Name         string `json:"name"`
			Slug         string `json:"slug"`
			Symbol       string `json:"symbol"`
			TokenAddress string `json:"token_address"`
		} `json:"platform"`
		Quote struct {
			Usd struct {
				LastUpdated      string  `json:"last_updated"`
				MarketCap        float64 `json:"market_cap"`
				PercentChange1h  float64 `json:"percent_change_1h"`
				PercentChange24h float64 `json:"percent_change_24h"`
				PercentChange30d float64 `json:"percent_change_30d"`
				PercentChange7d  float64 `json:"percent_change_7d"`
				Price            float64 `json:"price"`
				Volume24h        float64 `json:"volume_24h"`
			} `json:"USD"`
		} `json:"quote"`
		Slug        string   `json:"slug"`
		Symbol      string   `json:"symbol"`
		Tags        []string `json:"tags"`
		TotalSupply int64    `json:"total_supply"`
	} `json:"data"`
	Status struct {
		CreditCount  int64       `json:"credit_count"`
		Elapsed      int64       `json:"elapsed"`
		ErrorCode    int64       `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Notice       interface{} `json:"notice"`
		Timestamp    string      `json:"timestamp"`
		TotalCount   int64       `json:"total_count"`
	} `json:"status"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print("Errored: ", err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", loadenv.LoadEnv())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Errored while sending the request to the server: ", err)
		os.Exit(1)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var crypto Crypto
	json.Unmarshal([]byte(respBody), &crypto)

	for i := 0; i < len(crypto.Data); i++ {
		fmt.Println(crypto.Data[i].Name)
	}
}
