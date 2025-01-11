package ezgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCause(t *testing.T) {
	scope := NewScopeWithDefaultLogger("TestScheduler")

	// This test is not yet implemented.
	err := NewCause(origin(), "origin")
	LogCausesf(scope.GetLogger(), err, "TestCauses")
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
