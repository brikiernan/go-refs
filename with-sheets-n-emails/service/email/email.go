package email

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strings"
	"with-sheets-n-emails/types"
)

type EmailRequest struct {
	addr   string
	auth   smtp.Auth
	params types.EmailParams
}

func NewEmailRequest(addr string, auth smtp.Auth, params types.EmailParams) *EmailRequest {
	return &EmailRequest{
		addr:   addr,
		auth:   auth,
		params: params,
	}
}

func (r *EmailRequest) ParseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.params.Body = buf.String()
	return nil
}

func (r *EmailRequest) SendEmail() (bool, error) {
	from := "From: " + r.params.FromName + " <" + r.params.FromEmail + "> \n"
	to := "To: " + strings.Join(r.params.To, ", ") + "\n"
	replyTo := "Reply-To: " + r.params.ReplyTo + "\n"
	subject := "Subject: " + r.params.Subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(from + to + replyTo + subject + mime + "\n" + r.params.Body)

	err := smtp.SendMail(r.addr, r.auth, r.params.FromEmail, r.params.To, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}
