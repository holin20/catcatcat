package monitor

import (
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const (
	gmailSender = "catcatcattm@gmail.com"
	gmailPwEnv  = "secrets/google_app_password.txt"

	recipientEmail = "holin20@gmail.com"
)

type Notifier struct {
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyEmail(
	ruleName string,
	evalTime time.Time,
	result float64,
	resultTime time.Time,
) error {
	subject := "You've got a cat!"
	body := fmt.Sprintf(
		"Cat '%s' is just detected at %s (data obtained at %s / delay by %s)",
		ruleName,
		evalTime.Format(time.RFC1123),
		resultTime.Format(time.RFC1123),
		evalTime.Sub(resultTime).String(),
	)

	return ezgo.GmailSender().
		From(gmailSender).
		To(recipientEmail).
		Subject(subject).
		Body(body).
		PasswordEnv(gmailPwEnv).
		Send()
}
