package sampleapps

import (
	"net/smtp"
	"with-sheets-n-emails/service/email"
	"with-sheets-n-emails/types"
	"with-sheets-n-emails/utils"
)

type Emails struct {
	smtpHost string
	smtpAuth smtp.Auth
}

func NewEmails(smtpHost string, smtpAuth smtp.Auth) *Emails {
	return &Emails{
		smtpHost: smtpHost,
		smtpAuth: smtpAuth,
	}
}

func (e *Emails) SendConfirmationEmail(payload types.AppsRequestParams) (bool, error) {
	params := types.EmailParams{
		FromName:  utils.Getenv("FROM_NAME"),
		FromEmail: utils.Getenv("FROM_EMAIL"),
		To:        []string{payload.Email},
		Subject:   "We Recieved Your Request",
	}

	data := struct{ Name, Message, Email string }{
		Name:    payload.First,
		Message: "This is a test message in a HTML template test confirmation 2",
		Email:   payload.Email,
	}

	r := email.NewEmailRequest(e.smtpHost, e.smtpAuth, params)

	err := r.ParseTemplate("templates/email/sample-apps-confirmation.html", data)
	if err != nil {
		return false, err
	}

	return r.SendEmail()
}

func (e *Emails) SendNotificationEmail(payload types.AppsRequestParams) (bool, error) {
	fromEmail := utils.Getenv("FROM_EMAIL")
	params := types.EmailParams{
		FromName:  utils.Getenv("FROM_NAME"),
		FromEmail: fromEmail,
		To:        []string{fromEmail},
		ReplyTo:   payload.First + " " + payload.Last + " <" + payload.Email + ">",
		Subject:   "New Sample Apps Access Request",
	}

	data := struct{ Name, Message, Email string }{
		Name:    payload.First,
		Message: "This is a test message in a HTML template test notification 2",
		Email:   payload.Email,
	}

	r := email.NewEmailRequest(e.smtpHost, e.smtpAuth, params)

	err := r.ParseTemplate("templates/email/sample-apps-notification.html", data)
	if err != nil {
		return false, err
	}

	return r.SendEmail()
}
