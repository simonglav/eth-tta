package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalAmount(t *testing.T) {
	block_number := int(11509797)
	transactions := 155
	var amount float64 = 2.285404805647828

	Jresult, _ := GetTotalAmount(block_number)
	var tta_test tta
	json.Unmarshal(Jresult, &tta_test)
	assert.Equal(t, amount, tta_test.Amount)
	assert.Equal(t, transactions, tta_test.Transactions)
}
