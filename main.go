package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {

	mailingList := []string{} // TODO Load list from json

	send("OwO", mailingList)
}

func send(body string, to []string) {
	from := os.Getenv("MAIL_ID")
	pass := os.Getenv("MAIL_PASSWORD")

	msg := "Subject: Hello there o/ \n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent")
}
