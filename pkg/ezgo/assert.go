package ezgo

import (
	"fmt"
)

func AssertNotNil[T any](p *T, msg string) {
	if p != nil {
		return
	}
	Fatal(msg)
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
