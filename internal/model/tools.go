package model

import (
	"log"
	"math/big"
	"strconv"

	"github.com/pelletier/go-toml"
)

// WeiToEther is a Etherium denomination converter, from Wei to Ether
func WeiToEther(w *big.Int) float64 {
	if w.Sign() > 1 {
		return float64(0)
	}
	wei := new(big.Float)
	wei.SetInt(w)

	ether := new(big.Float).SetFloat64(1000000000000000000) // 10^18
	wei = wei.Quo(wei, ether)
	weis, _ := wei.Float64()
	return weis
}

// BuildURL is building etherscan.io API's endpoint with ETH block number and with API token from 'token.toml'
func BuildURL(block_number int) string {
	hexBlock := "0x" + strconv.FormatInt(int64(block_number), 16)

	APIKey := "YourApiKeyToken" // default token

	config, err := toml.LoadFile("token.toml")
	if err != nil {
		log.Println(err)
	} else {
		APIKey = config.Get("EtherScanAPIKey").(string)
	}

	return "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=" +
		hexBlock + "&boolean=true&apikey=" + APIKey
}
