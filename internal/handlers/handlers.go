package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/total-transactions-amount-eth/internal/model"
)

// ETHBlockTotal handles "/api/block/{block_number:[0-9]+}/total" GET request;
// Writes relevant http.Status in Header;
// Writes JSON {"transactions": int,"amount":float64} with zeros if any error occurs(with logging);
// Number of transactions for ETH block with given block_number and total transactions amount in Ethers
func ETHBlockTotal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	block, err := strconv.Atoi(params["block_number"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var totalAmount model.TotTransAm
		err := totalAmount.GetBlockTTA(block)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		JtotalAmount, err := json.Marshal(totalAmount)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(JtotalAmount)
	}
}
