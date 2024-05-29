package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
	smtpUser = "daldenuteryag@gmail.com"
	smtpPass = "kpzq rnyd wmtr kfiv"
)

// SendReceipt sends the receipt to the customer's email.
func SendReceipt(toEmail, receiptPath string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your Purchase Receipt")
	m.SetBody("text/plain", "Thank you for your purchase. Please find your receipt attached.")
	m.Attach(receiptPath)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}
	return nil
}
