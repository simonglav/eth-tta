package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/pelletier/go-toml"
)

// Dafault JSON in case if any error ocurrs
var EmptyJSON []byte = []byte(`{"transactions":0,"amount":0}`)

// Raw ETH block data
type block struct {
	Result struct {
		Transactions []struct {
			Value string `json:"value"`
		} `json:"transactions"`
	}
}

// GetTotalAmount returns JSON []byte of transactions number and a sum of the values;
// {"transactions": int,"amount":float64};
// If error occurs returns EmptyJSON and error
func GetTotalAmount(block_number int) ([]byte, error) {
	hexBlock := "0x" + strconv.FormatInt(int64(block_number), 16)

	resp, err := http.Get(buildURL(hexBlock))
	if err != nil {
		return EmptyJSON, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmptyJSON, err
	}

	var blc block
	if err := json.Unmarshal(body, &blc); err != nil {
		return EmptyJSON, err
	}

	if blc.Result.Transactions == nil {
		return EmptyJSON, errors.New("invalid ETH block number")
	}

	var result tta
	result.calculate(&blc)

	Jresult, err := json.Marshal(result)
	if err != nil {
		return EmptyJSON, err
	}
	return Jresult, nil
}

// buildURL is building API's URL with hexBlock as ETH block number(0x format) and with API token from 'token.toml'
func buildURL(hexBlock string) string {
	APIKey := "YourApiKeyToken" // default token

	config, err := toml.LoadFile("token.toml")
	if err != nil {
		log.Println(err)
	} else {
		APIKey = config.Get("EtherScanAPIKey").(string)
	}

	return "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=" + hexBlock + "&boolean=true&apikey=" + APIKey
}
