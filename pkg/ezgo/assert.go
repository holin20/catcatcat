package ezgo

import (
	"fmt"
)

func Assert(cond bool, msg string) {
	if cond {
		return
	}
	Fatal(msg)
}

func Assertf(cond bool, msgFmt string, args ...any) {
	if cond {
		return
	}
	Fatal(fmt.Sprintf(msgFmt, args...))
}

func AssertAlways(msg string) {
	Assert(false, msg)
}

func AssertAlwaysf(msgFmt string, args ...any) {
	AssertAlways(fmt.Sprintf(msgFmt, args...))
}

func AssertNotNil[T any](p *T, msg string) {
	Assert(p != nil, msg)
}

func AssertNotNilf[T any](p *T, msgFmt string, args ...any) {
	AssertNotNil(p, fmt.Sprintf(msgFmt, args...))
}

func AssertNoError(err error, msg string) {
	if err == nil {
		return
	}
	Fatal(NewCause(err, msg).Error())
}

func AssertNoErrorf(err error, msgFmt string, args ...any) {
	AssertNoError(err, fmt.Sprintf(msgFmt, args...))
}

func Fatal(msg string) {
	fmt.Println(msg)
	panic(msg)
}
