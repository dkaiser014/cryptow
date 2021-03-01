package getdata

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"cryptow/pkg/loadenv"
)

// GetData returns the latest data of the top 10
// cryptocurrencies using the Coinmarketcap API
func GetData() []byte {
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
	return respBody
}
