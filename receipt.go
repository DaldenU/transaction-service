package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func GetTransaction(transactionID string) (Transaction, error) {
	var transaction Transaction
	row := db.QueryRow("SELECT id, customer_id, customer_name, customer_email, status FROM transactions WHERE id = ?", transactionID)
	err := row.Scan(&transaction.ID, &transaction.CustomerID, &transaction.CustomerName, &transaction.CustomerEmail, &transaction.Status)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GenerateReceipt(transaction Transaction, transactionRequest TransactionRequest) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Receipt")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Beauty Salon")
	pdf.Ln(10)
	pdf.Cell(40, 10, "=============================================")
	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Transaction ID: %s", transaction.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Customer Name: %s", transaction.CustomerName))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Customer Email: %s", transaction.CustomerEmail))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", time.Now().Format(time.RFC1123)))
	pdf.Ln(10)
	pdf.Cell(40, 10, "=============================================")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Purchased Items:")
	pdf.Ln(10)
	var totalAmount float64
	for _, item := range transactionRequest.CartItems {
		itemTotal := item.Price * item.Quantity
		totalAmount += itemTotal
		pdf.Cell(40, 10, fmt.Sprintf("Item: %s", item.Name))
		pdf.Ln(10)
		pdf.Cell(40, 10, fmt.Sprintf("%.2f x %.2f = Total: %.2f", item.Price, item.Quantity, itemTotal))
		pdf.Ln(10)
	}
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "=============================================")
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Total Amount: %.2f", totalAmount))
	pdf.Ln(10)
	pdf.Cell(40, 10, "Payed with Bank Card")

	receiptsDir := "./receipts"
	if _, err := os.Stat(receiptsDir); os.IsNotExist(err) {
		log.Printf("Directory %s does not exist, creating...\n", receiptsDir)
		err := os.Mkdir(receiptsDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create receipts directory: %v", err)
		}
	}

	receiptPath := fmt.Sprintf("%s/receipt_%s.pdf", receiptsDir, transaction.ID)
	err := pdf.OutputFileAndClose(receiptPath)
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %v", err)
	}

	log.Printf("PDF generated successfully at %s\n", receiptPath)
	return receiptPath, nil
}
