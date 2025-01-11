package ezgo

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type cause struct {
	msg string

	// A cause is either due to a "direct cause" or an "error". The latter implies this is the root cause.
	directCause *cause
	err         error
}

func NewCausef(err error, msgFmt string, args ...any) *cause {
	return NewCause(err, fmt.Sprintf(msgFmt, args...))
}

func NewCause(err error, msg string) *cause {
	if err == nil {
		return nil
	}

	if c, ok := err.(*cause); ok && c != nil {
		return &cause{
			msg:         msg,
			directCause: c,
		}
	}
	return &cause{
		msg: msg,
		err: err,
	}
}

func (c *cause) String() string {
	return c.Error()
}

func (c *cause) Error() string {
	return c.buildRootCausingString()
}

func (c *cause) IsRootCause() bool {
	return c.err != nil
}

func (c *cause) GetRootCause() *cause {
	for c != nil {
		if c.IsRootCause() {
			return c
		}
		c = c.directCause
	}
	return nil
}

func (c *cause) Traceback() ([]*cause, *cause) {
	var causes []*cause
	var rootCause *cause
	for c != nil {
		causes = append(causes, c)
		if c.IsRootCause() {
			rootCause = c
			break
		}
		c = c.directCause
	}
	return causes, rootCause
}

func (c *cause) buildRootCausingString() string {
	causes, rootCause := c.Traceback()
	var causesStr []string
	for _, cause := range causes {
		causesStr = append(causesStr, cause.msg)
	}
	return fmt.Sprintf(
		"Error: %s | Root Cause: %s | Causes: %s",
		causes[0].msg,
		IfLazy(
			rootCause.err != nil,
			func() string { return rootCause.err.Error() },
			func() string { return "[No Root Cause]" },
		),
		strings.Join(causesStr, " <- "),
	)
}

func LogCauses(logger *zap.Logger, err error, msg string) {
	causes, rootCause := NewCause(err, msg).Traceback()
	var causesStr []string
	for _, cause := range causes {
		causesStr = append(causesStr, cause.msg)
	}

	logger.Error(causes[0].msg, zap.Error(rootCause.err), zap.String("backtrace", strings.Join(causesStr, " <- ")))
}

func LogCausesf(logger *zap.Logger, err error, msgFmt string, args ...any) {
	LogCauses(logger, err, fmt.Sprintf(msgFmt, args...))
}
