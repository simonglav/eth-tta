package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalAmount(t *testing.T) {
	block_number := int(11509797)
	transactions := 155
	var amount float64 = 2.285404805647828
	var tta TotTransAm
	tta.GetBlockTTA(block_number)
	assert.Equal(t, amount, tta.Amount)
	assert.Equal(t, transactions, tta.Transactions)
}
