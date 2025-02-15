package monitor

import (
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const (
	gmailSender = "catcatcattm@gmail.com"
	gmailPwEnv  = "CATCATCAT_GAPP_PW"

	recipientEmail = "holin20@gmail.com"
)

type Notifier struct {
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyEmail(
	ruleName string,
	cdpRuleConfig *CdpRuleConfig,
	evalTime time.Time,
	result float64,
	resultTime time.Time,
) error {
	subject := "You've got a cat! " + ruleName

	currentStatus := fmt.Sprintf(cdpRuleConfig.QueryResultTemplate, result)

	body := fmt.Sprintf(
		"%s was just detected to meet your watch criteria at %s (data point from %s ago)\n"+
			"- Alert Criteria: %s\n"+
			"- Current Status: %s\n",
		ruleName,
		evalTime.Format(time.RFC1123),
		evalTime.Sub(resultTime).String(),
		cdpRuleConfig.AlertCriteria,
		currentStatus,
	)

	return ezgo.GmailSender().
		From(gmailSender).
		To(recipientEmail).
		Subject(subject).
		Body(body).
		PasswordEnv(gmailPwEnv).
		Send()
}
