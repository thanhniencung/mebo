package mail

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
	"sync"
)

type Mail struct {
	auth smtp.Auth
	from string
	pass string
	mailServer string
}

var instance *Mail
var once sync.Once

// Setup some info for mail package (we need to connect to mail server first)
func Init(from, pass, mailServer string) {
	once.Do(func() {
		instance = &Mail{
			auth: smtp.PlainAuth("", from, pass, mailServer),
			from: from,
			pass: pass,
			mailServer: mailServer,
		}
	})
}

func New() *Mail {
	return instance
}

// This func going to send an email with html template
// if err = 535 5.7.8 Username and Password not accepted. Learn more at\n5.7.8  https://support.google.com/mail/?p=BadCredentials 10sm6673282pfh.14 - gsmtp
// go to link >> https://myaccount.google.com/lesssecureapps?pli=1 and turn on
func (m *Mail) Send(to, subject, templatePath string, data interface{}) error {
	if len(to) == 0 || len(subject) == 0 || len(templatePath) == 0 {
		return errors.New("to, subject, templatePath can not empty")
	}

	body, err := m.parseTemplate(templatePath, data)
	if err != nil {
		return err
	}

	var msgs []string

	msgs = append(msgs, "From: Mebo<" + m.from + ">\r")
	msgs = append(msgs, "To: " + to + "\r")
	msgs = append(msgs, "Subject: " + subject + "\r")
	msgs = append(msgs, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n")
	msgs = append(msgs, body + "\r")

	msg := []byte (strings.Join(msgs, "\n"))
	mailPort := fmt.Sprintf("%s:%d", m.mailServer, 587)

	return smtp.SendMail(mailPort, m.auth, m.from, []string{to}, msg)
}

// Helper function help you parse tempate and binding data to template
func (e *Mail) parseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

