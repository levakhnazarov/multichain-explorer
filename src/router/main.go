package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/CheckBuyerTx", CheckBuyerTx).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
