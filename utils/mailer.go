package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
)

type SignatureMailPayload struct {
	RequesterEmail string
	RequesterName  string
	QrCodeToken    string
}

func SendSignatureMail(payload SignatureMailPayload) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	//Temporarily Create QR Code file to send later
	qrCodeUrl := "https://chart.apis.google.com/chart?cht=qr&chs=300x300&chl="
	qrCodeUrl += payload.QrCodeToken
	err = DownloadFile("signature.png", qrCodeUrl)
	if err != nil {
		panic(err)
	}

	//Configure Email Properties
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", payload.RequesterEmail)
	mailer.SetHeader("Subject", "Test mail")
	emailBody := fmt.Sprintf("Hello, <b>%s</b> <br> Your signature request has been accepted", payload.RequesterName)
	mailer.SetBody("text/html", emailBody)
	mailer.Attach("./signature.png")

	CONFIG_SMTP_PORT, err := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	if err != nil {
		log.Println("Error fetching env for mailer")
	}
	// Prepare dialer as an actor to send email
	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		CONFIG_SMTP_PORT,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Remove("signature.png")
	if err != nil {
		log.Printf("Error Removing Temp File!")
	}

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
