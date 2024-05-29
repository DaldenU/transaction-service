package main

import (
	"database/sql"
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
    `
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveTransaction(req TransactionRequest) (string, error) {
	stmt, err := db.Prepare("INSERT INTO transactions(customer_id, customer_name, customer_email, status) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(req.Customer.ID, req.Customer.Name, req.Customer.Email, "awaiting payment")
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func UpdateTransactionStatus(transactionID, status string) error {
	_, err := db.Exec("UPDATE transactions SET status = ? WHERE id = ?", status, transactionID)
	return err
}
