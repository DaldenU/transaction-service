package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/static/"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	})
	router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/create.html")
	})
	router.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/payment.html")
	})
	router.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/transactions.html")
	})
	router.HandleFunc("/create-transaction", CreateTransactionHandler).Methods("POST")
	router.HandleFunc("/process-payment", ProcessPaymentHandler).Methods("POST")
	router.HandleFunc("/transactions/{customerId}", CustomerTransactionsHandler).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
