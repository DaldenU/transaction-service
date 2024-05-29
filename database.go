package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./transactions.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS transactions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer_id TEXT,
        customer_name TEXT,
        customer_email TEXT,
        status TEXT
    );
    CREATE TABLE IF NOT EXISTS transaction_requests (
        transaction_id INTEGER PRIMARY KEY,
        cart_items TEXT,
        customer TEXT
    );
    `
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveTransaction(req TransactionRequest) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}

	stmt, err := tx.Prepare("INSERT INTO transactions(customer_id, customer_name, customer_email, status) VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	res, err := stmt.Exec(req.Customer.ID, req.Customer.Name, req.Customer.Email, "awaiting payment")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	cartItems, err := json.Marshal(req.CartItems)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	customer, err := json.Marshal(req.Customer)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	stmt, err = tx.Prepare("INSERT INTO transaction_requests(transaction_id, cart_items, customer) VALUES(?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = stmt.Exec(id, cartItems, customer)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func UpdateTransactionStatus(transactionID, status string) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions WHERE id = ?", transactionID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("transaction with ID %s not found", transactionID)
	}

	_, err = db.Exec("UPDATE transactions SET status = ? WHERE id = ?", status, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func GetTransactionRequest(transactionID string) (TransactionRequest, error) {
	var transactionRequest TransactionRequest
	var cartItems, customer string

	err := db.QueryRow("SELECT cart_items, customer FROM transaction_requests WHERE transaction_id = ?", transactionID).Scan(&cartItems, &customer)
	if err != nil {
		return transactionRequest, err
	}

	err = json.Unmarshal([]byte(cartItems), &transactionRequest.CartItems)
	if err != nil {
		return transactionRequest, err
	}
	err = json.Unmarshal([]byte(customer), &transactionRequest.Customer)
	if err != nil {
		return transactionRequest, err
	}

	return transactionRequest, nil
}
