package libs

import (
	"os"
	"strconv"

	"github.com/pius706975/backend/database/models"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL      string
	Username string
	Subject  string
}

func SendEmail(user *models.User, data *EmailData) error {

	h := hermes.Hermes{

		Product: hermes.Product{
			Name: "Online Notes",
			Link: os.Getenv("BASE_URL"),
			Logo: "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_LightBlue.png",
		},
	}

	emailBody, err := h.GenerateHTML(hermes.Email{

		Body: hermes.Body{

			Name: user.Username,
			Intros: []string{
				"This is a verification step.",
			},

			Actions: []hermes.Action{
				{
					Instructions: "Click on the button below to verify your email immediately.",
					Button: hermes.Button{
						Color: "#0D5365",
						Text:  "Click to verify",
						Link:  data.URL,
					},
				},
			},
		},
	})

	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "Onine Notes <example@gmail.com>")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", emailBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD")) 

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
