package main

// PaymentForm represents the form for entering payment details.
type PaymentForm struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CVV            string `json:"cvv"`
	Name           string `json:"name"`
	Address        string `json:"address"`
}

// PaymentRequest represents a request to process a payment.
type PaymentRequest struct {
	TransactionID string      `json:"transactionID"`
	PaymentForm   PaymentForm `json:"paymentForm"`
}
