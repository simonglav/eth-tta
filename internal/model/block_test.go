package model

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeiToEther(t *testing.T) {
	bgnmb1, _ := new(big.Int).SetString("2285404805647828000", 10)
	bgnmb2, _ := new(big.Int).SetString("1130987085446826418823", 10)
	bgnmb3, _ := new(big.Int).SetString("4998770090000000000", 10)
	testCases := []struct {
		name    string
		wei     *big.Int
		ether   float64
		isValid bool
	}{
		{
			name:    "first number valid",
			wei:     bgnmb1,
			ether:   2.285404805647828,
			isValid: true,
		},
		{
			name:    "second number valid",
			wei:     bgnmb2,
			ether:   1130.9870854468263,
			isValid: true,
		},
		{
			name:    "third number valid",
			wei:     bgnmb3,
			ether:   4.99877009,
			isValid: true,
		},
		{
			name:    "invalid calculations",
			wei:     bgnmb1,
			ether:   123.456,
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.Equal(t, tc.ether, WeiToEther(tc.wei))
			}
		})
	}
}
