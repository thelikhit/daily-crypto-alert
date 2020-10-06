package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

type Coin struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Price  string `json:"price"`
}

func main() {

	// HTTP GET request to api.nomics.com using demo API key
	response, err := http.Get("https://api.nomics.com/v1/currencies/ticker?key=demo-26240835858194712a4f8cc0dc635c7a&ids=BTC,ETH,VET,ADA&convert=USD")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var coin[] Coin
	err = json.Unmarshal(responseData, &coin)
	if err != nil {
		log.Fatal(err)
	}

	var body strings.Builder
	body.WriteString("Hello!\n\nHere are the prices of cryptocurrencies today:\n")
	for _, value := range coin {
		body.WriteString("\n")
		body.WriteString(value.Name)
		body.WriteString(" ")
		body.WriteString(value.Symbol)
		body.WriteString(" $")
		body.WriteString(value.Price)
	}

	from := "...@gmail.com"
	password := "<password>"

	to := []string{
		"...@gmail.com",
		"...@gmail.com",
	}

	// SMTP Server Config.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Daily Crypto Price Alert\n\n" + body.String())

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Success message
	fmt.Println("Email Sent Successfully!")
}
