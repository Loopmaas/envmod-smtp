package smtp

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// ================================================================
//
// ================================================================
type Smtp struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DisplayName  string
	DisplayEmail string
}

func New() (*Smtp, error) {
	return &Smtp{
		Host:         os.Getenv("SMTP_HOST"),
		Port:         os.Getenv("SMTP_PORT"),
		Username:     os.Getenv("SMTP_USERNAME"),
		Password:     os.Getenv("SMTP_PASSWORD"),
		DisplayName:  os.Getenv("SMTP_DISPLAY_NAME"),
		DisplayEmail: os.Getenv("SMTP_DISPLAY_EMAIL"),
	}, nil
}

// ================================================================
//
// ================================================================
func (e Smtp) SendHTMLEmail(to []string, bcc []string, subject, body string) error {
	server := e.Host + ":" + e.Port
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	recipients := to
	if len(bcc) > 0 {
		recipients = append(to, bcc...)
	}

	msgs := []string{
		`MIME-version: 1.0;`,
		`Content-Type: text/html; charset="UTF-8";`,
		fmt.Sprintf("From: %s <%s>", e.DisplayName, e.DisplayEmail),
		fmt.Sprintf("To: %s", strings.Join(to, ", ")),
		fmt.Sprintf("Subject: %s", subject),
		body,
	}

	return smtp.SendMail(server, auth, e.Username, recipients, []byte(strings.Join(msgs, "\n")))
}
