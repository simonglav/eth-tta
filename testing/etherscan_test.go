package testing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EthersanAPIAccessebilityTest(t *testing.T) {
	URL := "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=0xafa01b&boolean=true&apikey=YourApiKeyToken"
	resp, err := http.Get(URL)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	defer resp.Body.Close()
}
