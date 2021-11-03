package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
)

// Raw ETH block data
type block struct {
	Result struct {
		Transactions []struct {
			Value string `json:"value"`
		} `json:"transactions"`
	}
}

// Total Transactions Amount
type TotTransAm struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}

// GetBlockTTA is computing number of transactions for ETH block and total transactions amount in Ethers
func (tta *TotTransAm) GetBlockTTA(block_number int) error {
	resp, err := http.Get(BuildURL(block_number))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var blc block
	if err := json.Unmarshal(body, &blc); err != nil {
		return err
	}

	if blc.Result.Transactions == nil {
		return errors.New("invalid ETH block number")
	}

	tta.calculate(&blc)
	return nil

}

// calcualte is calculating total transactions number and a sum of the value of each transaction;
// *Amount is in Ether denomination
func (result *TotTransAm) calculate(blc *block) {
	totalAmount := new(big.Int)
	for _, trans := range blc.Result.Transactions {
		decimal := new(big.Int)
		decimal.SetString(trans.Value[2:], 16) // cut 0x and convert to decimal
		totalAmount.Add(totalAmount, decimal)
	}

	result.Transactions = len(blc.Result.Transactions)
	result.Amount = WeiToEther(totalAmount)
}
