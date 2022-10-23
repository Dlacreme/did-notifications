package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/dlacreme/gotifications/env"
)

type smtpConfig struct {
	username string
	addr     string
	auth     smtp.Auth
}

var smtpC smtpConfig

func Load() {
	smtpC = loadSMTPConfig()
}

func buildHeader(from string, to []string, subject string) string {
	fromHeader := fmt.Sprintf("From: %s", from)
	toHeader := "To: "
	for _, e := range to {
		toHeader += fmt.Sprintf("%s,", e)
	}
	toHeader = strings.TrimSuffix(toHeader, ",")
	subjectHeader := fmt.Sprintf("Subject: %s", subject)

	return fromHeader + "\r\n" + toHeader + "\r\n" + subjectHeader + "\r\n"
}

func buildMessage(from string, to []string, subject string, body string) []byte {
	header := buildHeader(from, to, subject)
	return []byte(header + "\r\n" + body + "\r\n")
}

func SendEmail(from string, to []string, subject string, body string) error {
	message := buildMessage(from, to, subject, body)
	if env.IsDebug() {
		fmt.Printf("Sending email through [%s]\n----\n%s\n----\n", smtpC.addr, message)
	}
	return smtp.SendMail(smtpC.addr, smtpC.auth, from, to, message)
}

func loadSMTPConfig() smtpConfig {
	port := os.Getenv("SMTP_PORT")
	hostname := os.Getenv("SMTP_HOSTNAME")
	addr := fmt.Sprintf("%s:%s", hostname, port)
	username := os.Getenv("SMTP_USERNAME")
	passwd := os.Getenv("SMTP_PASSWORD")
	fmt.Printf("PORT [%s]\nHOSTNAME [%s]\nADDR [%s]\n USERNAME [%s]\nPASSWD [%s]", port, hostname, addr, username, passwd)
	return smtpConfig{
		username: username,
		addr:     addr,
		auth:     smtp.PlainAuth("", username, passwd, hostname),
	}
}
