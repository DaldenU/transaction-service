package main

import "time"

type PaymentForm struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CVV            string `json:"cvv"`
	Name           string `json:"name"`
	Address        string `json:"address"`
}

type PaymentRequest struct {
	TransactionID string      `json:"transactionID"`
	PaymentForm   PaymentForm `json:"paymentForm"`
}

type Transaction struct {
	ID            string    `json:"id"`
	CustomerID    string    `json:"customerId"`
	CustomerName  string    `json:"customerName"`
	CustomerEmail string    `json:"customerEmail"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type CartItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TransactionRequest struct {
	CartItems []CartItem `json:"cartItems"`
	Customer  Customer   `json:"customer"`
}
