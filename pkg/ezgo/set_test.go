package ezgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := MakeSet(1, 2, 3, 4, 5)
	assert.False(t, s.Empty())
	assert.Equal(t, 5, s.Size())

	s.Remove(4, 5)
	assert.False(t, s.Empty())
	assert.Equal(t, 3, s.Size())

	subsetOfS := MakeSet(1, 2)
	assert.True(t, s.Covers(subsetOfS))
	assert.True(t, subsetOfS.CoveredBy(s))

	s = MakeSet(1, 2, 3, 4)
	p := MakeSet(2, 3, 4, 5)
	assert.False(t, s.Covers(p))
	assert.False(t, p.CoveredBy(s))

	sSubstractP := s.Substract(p)
	assert.Equal(t, sSubstractP.Size(), 1)
}
