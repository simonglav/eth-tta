package model

import (
	"math/big"
)

// Total Transactions Amount
type tta struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}

// calcualte is calculating total transactions number and a sum of the value of each transaction;
// *Amount is in Ether denomination
func (result *tta) calculate(blc *block) {
	totalAmount := new(big.Int)
	for _, trans := range blc.Result.Transactions {
		decimal := new(big.Int)
		decimal.SetString(trans.Value[2:], 16) // cut 0x and convert to decimal
		totalAmount.Add(totalAmount, decimal)
	}

	result.Transactions = len(blc.Result.Transactions)
	result.Amount = WeiToEther(totalAmount)
}

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
