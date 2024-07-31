package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"with-sheets/types"
)

// TODO refactor to use injection for emails
func SendSampleAppsRequestEmail(payload types.Request) {
	from := Getenv("FROM_EMAIL")
	password := Getenv("FROM_EMAIL_APP_PASSWORD")
	to := []string{payload.Email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("templates/sample-apps-request.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.Write([]byte(fmt.Sprintf("From: Add Name <%s>\r\n"+
		"To: %s\r\n"+
		// "Reply-To: example@email.com\r\n" +
		"Subject: Request Received\r\n"+mimeHeaders,
		from,
		payload.Email,
	),
	))

	t.Execute(&body, struct {
		Name, Message, Email string
	}{
		Name:    payload.Name,
		Message: "This is a test message in a HTML template",
		Email:   payload.Email,
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
