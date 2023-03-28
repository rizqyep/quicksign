package utils

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
	QrCodeUrl      string
}

type RejectedSignatureMailPayload struct {
	RequesterEmail string
	RequesterName  string
	Description    string
}

type ResetPasswordMailPayload struct {
	Email string
	Token string
}

func SendSignatureMail(payload SignatureMailPayload) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	//Temporarily Create QR Code file to send later

	err = DownloadFile("signature.png", payload.QrCodeUrl)
	if err != nil {
		panic(err)
	}

	//Configure Email Properties
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", payload.RequesterEmail)
	mailer.SetHeader("Subject", "Signature Request Accepted !")
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

func SendRejectedSignatureMail(payload RejectedSignatureMailPayload) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	//Configure Email Properties
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", payload.RequesterEmail)
	mailer.SetHeader("Subject", "Reset Your Password")
	emailBody := fmt.Sprintf("Hello, %s ! <br> We are sorry to inform that your signature request for '%s' has been rejected", payload.RequesterName, payload.Description)
	mailer.SetBody("text/html", emailBody)

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
}

func SendResetPasswordLink(payload ResetPasswordMailPayload) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	resetPasswordLink := os.Getenv("URL") + fmt.Sprintf("/reset-password/%s", payload.Token)
	//Configure Email Properties
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", payload.Email)
	mailer.SetHeader("Subject", "Reset Your Password")
	emailBody := fmt.Sprintf("Hello! <br> We have received a request to reset your password <br> Please kindly access this link : <a href='%s'>%s</a> to reset your password", resetPasswordLink, resetPasswordLink)
	mailer.SetBody("text/html", emailBody)

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
