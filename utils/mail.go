package utils

import (
	"log"
	"net/smtp"
	"os"
)

func SendMail(to, sub, message string) {
    from := os.Getenv("EMAIL")
    password := os.Getenv("MAILPASS")
 
    recipient := []string{
       to,
    }
 
    smtpHost := "smtp.mail.ru"
    smtpPort := "25"
 
    msg := []byte("Subject: " + sub + "\r\n" + message)
 
    auth := smtp.PlainAuth("", from, password, smtpHost)
 
    err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, recipient, msg)
    if err != nil {
       log.Println(err)
       return
    }
    
    log.Println("Warning message was send to " + to)
}