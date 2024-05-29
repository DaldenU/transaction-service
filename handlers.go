package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save transaction to database
	transactionID, err := SaveTransaction(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return transaction ID to the client
	w.Header().Set("Content-Type", "application/json")
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

		transactionRequest, err := GetTransactionRequest(req.TransactionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		receiptPath, err := GenerateReceipt(transaction, transactionRequest)
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

func RetrieveTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("customerId")
	if customerID == "" {
		http.Error(w, "customerId is required", http.StatusBadRequest)
		return
	}

	transactions, err := GetTransactionsByCustomerID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

func CustomerTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]

	if customerId == "" {
		http.Error(w, "customerId is required", http.StatusBadRequest)
		return
	}

	transactions, err := GetCustomerTransactions(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"transactions": transactions})
}
