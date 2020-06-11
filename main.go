package main

import (
	"log"
	"net/smtp"
	"os"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/urfave/cli"
)

const (
	from = "kholikov.x@gmail.com"
	pass = "Hamid123456"
	to   = "kholikov.x@gmail.com"


	accountSid = "AC1503e08dd56cf5209b3503b03e7c1d1c"
	authToken = "73c8cda551bbfb7ce3d3c34557efb6d3"

	fromPhone = "+12139479583"
	toPhone   = "+998901233323"
)

func sendToMail(body string) {

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Assolomu alekum !!!\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func smsToPhone(body string) {
	// Set account keys & information
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", toPhone)
	msgData.Set("From", fromPhone)
	msgData.Set("Body", body )
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

func main() {

	informationBody := "Your purchases \nProduct: \n   Lavash \n Price: 2$ "

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name: "sms",
			Action: func(c *cli.Context) error {
				smsToPhone(informationBody)
				return nil
			},
		},
		{
			Name: "mail",
			Action: func(c *cli.Context) error {
				sendToMail(informationBody)
				return nil
			},
		},
	}
	app.Run(os.Args)

}
