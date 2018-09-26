package main

import (
	"encoding/json"
	"explorers"
	"model"
	"net/http"
	"storage"
	"strconv"
	"fmt"
)

/**
params: amount timestamp sender receiver ticker

*/

func CheckBuyerTx(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var query = r.URL.Query()
	sendedAmount := query.Get("received_amount")

	tx := &model.Tx{Timestamp: query.Get("timestamp"), Sender: query.Get("sender"), Receiver: query.Get("receiver")}
	amount, err := strconv.ParseFloat(sendedAmount, 64)
	if err != nil {
		json.NewEncoder(w).Encode(&model.JsonResponse{
			Success: false,
			Comment: err.Error(),
		})
		return
	}
	tx.Amount = amount

	if query.Get("ticker") == "" || tx.Sender == "" || tx.Receiver == "" || sendedAmount == "" {
		json.NewEncoder(w).Encode(&model.JsonResponse{
			Success: false,
			Comment: "please provide required parameters (ticker, senderAddress, receiverAddress, sendedAmount)",
		})
		return
	}

	coin, err := storage.GetSingleCoinExplorer(query.Get("ticker"))
	if err != nil {
		json.NewEncoder(w).Encode(&model.JsonResponse{
			Success: false,
			Comment: err.Error(),
		})
		return
	}
	tx.Explorer = coin.Explorer
	tx.Ticker = coin.Ticker

	if err = explorers.NewAggregator().CheckIfTxExist(tx);err == nil {
		json.NewEncoder(w).Encode(&model.JsonResponse{
			Success: true,
		})
		return
	} else {
		fmt.Println(err.Error())


		json.NewEncoder(w).Encode(&model.JsonResponse{
			Success: false,
		})
		return
	}

}
