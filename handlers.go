package main

import (
	"encoding/json"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save transaction to database (placeholder function)
	transactionID, err := SaveTransaction(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return transaction ID to the client
	json.NewEncoder(w).Encode(map[string]string{"transactionID": transactionID})
}

func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulate successful payment
	success := true

	if success {
		err := UpdateTransactionStatus(req.TransactionID, "paid")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		transaction, err := GetTransaction(req.TransactionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		receiptPath, err := GenerateReceipt(transaction, TransactionRequest{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = SendReceipt(transaction.CustomerEmail, receiptPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
