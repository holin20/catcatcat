package ezgo

import (
	"fmt"
	"strings"
)

type cause struct {
	msg string

	// A cause is either due to a "direct cause" or an "error". The latter implies this is the root cause.
	directCause *cause
	err         error
}

func NewCause(err error, msgFmt string, args ...any) *cause {
	if c, ok := err.(*cause); ok && c != nil {
		return &cause{
			msg:         fmt.Sprintf(msgFmt, args...),
			directCause: c,
		}
	}
	return &cause{
		msg: fmt.Sprintf(msgFmt, args...),
		err: err,
	}
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
		strings.Join(causesStr[1:], " <- "),
	)
}

func PrintCauses(err error, msgFmt string, args ...any) {
	fmt.Println(NewCause(err, msgFmt, args...))
}
