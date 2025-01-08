package ezgo

import (
	"fmt"
	"net/smtp"
	"os"
)

type gmailSender struct {
	from    string
	to      string
	subject string
	body    string
	pwEnv   string
}

func GmailSender() *gmailSender {
	return &gmailSender{}
}

func (g *gmailSender) From(from string) *gmailSender {
	g.from = from
	return g
}

func (g *gmailSender) To(to string) *gmailSender {
	g.to = to
	return g
}

func (g *gmailSender) Body(body string) *gmailSender {
	g.body = body
	return g
}

func (g *gmailSender) PasswordEnv(pwEnv string) *gmailSender {
	g.pwEnv = pwEnv
	return g
}

func (g *gmailSender) Subject(subject string) *gmailSender {
	g.subject = subject
	return g
}

func (g *gmailSender) Send() error {
	msgFmt :=
		`From: %s
To: %s
Subject: %s

%s`

	// Arguments sanity checks
	if g.from == "" {
		return fmt.Errorf("from is not set")
	}
	if g.to == "" {
		return fmt.Errorf("to is not set")
	}
	if g.pwEnv == "" {
		return fmt.Errorf("pwEnv is not set")
	}
	pw := os.Getenv(g.pwEnv)
	if pw == "" {
		return fmt.Errorf("password is not set in env %s", g.pwEnv)
	}

	msg := fmt.Sprintf(msgFmt, g.from, g.to, g.subject, g.body)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", g.from, pw, "smtp.gmail.com"),
		g.from,
		[]string{g.to},
		[]byte(msg),
	)
}
