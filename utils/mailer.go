package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-gomail/gomail"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "jongdigitalid@gmail.com"
const CONFIG_AUTH_EMAIL = "jongdigitalid@gmail.com"
const CONFIG_AUTH_PASSWORD = "wggb hkzm glnv rttf"

func SendSignatureMail() {

	qrCodeUrl := "https://chart.apis.google.com/chart?cht=qr&chs=300x300&chl="

	qrCodeUrl += "abccdderff"

	err := DownloadFile("qrcode.png", qrCodeUrl)
	if err != nil {
		panic(err)
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "rizqyepr@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("./qrcode.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Remove("qrcode.png")
	log.Println("Mail sent!")

}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
