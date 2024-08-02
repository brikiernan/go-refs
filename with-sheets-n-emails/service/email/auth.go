package email

import (
	"net/smtp"
	"with-sheets-n-emails/utils"
)

func Auth() (string, smtp.Auth) {
	from := utils.Getenv("FROM_EMAIL")
	password := utils.Getenv("FROM_EMAIL_APP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	addr := smtpHost + ":" + smtpPort
	auth := smtp.PlainAuth("", from, password, smtpHost)
	return addr, auth
}
