package ezgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gScope *Scope = Must(NewScopeWithDefaultLogger())

func TestCause(t *testing.T) {
	// This test is not yet implemented.
	err := NewCause(origin(), "origin")
	LogCauses(gScope.GetLogger(), err, "TestCauses")
	assert.False(t, true)
}

func root() error {
	return fmt.Errorf("just not feeling well")
}

func origin() error {
	return NewCause(intermidiate(), "intermidiate")
}

func intermidiate() error {
	return NewCause(root(), "root")
}
