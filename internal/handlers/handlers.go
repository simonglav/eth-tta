package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/total-transactions-amount-eth/internal/model"
	"github.com/total-transactions-amount-eth/internal/store"
)

// ETHBlockTotal handles "/api/block/{block_number:[0-9]+}/total" GET request;
// Writes relevant http.Status in Header;
// Writes JSON {"transactions": int,"amount":float64} with zeros if any error occurs(with logging);
// Number of transactions for ETH block with given block_number and total transactions amount in Ethers
func ETHBlockTotal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	block_number := params["block_number"]
	w.Header().Set("Content-Type", "application/json")
	cachedTotalAmount, err := store.GetCache(block_number)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(cachedTotalAmount)
		return
	}
	block, err := strconv.Atoi(block_number)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var totalAmount model.TotTransAm
	err = totalAmount.GetBlockTTA(block)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jTotalAmount, err := json.Marshal(totalAmount)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.SetCache(block_number, jTotalAmount)
	w.WriteHeader(http.StatusCreated)
	w.Write(jTotalAmount)
}
