package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	blockNumber := "1234"
	jsonRes := []byte(`{"transactions":123,"amount":12345}`)
	SetCache(blockNumber, jsonRes)
	recived, _ := GetCache(blockNumber)
	assert.Equal(t, jsonRes, recived)
}
