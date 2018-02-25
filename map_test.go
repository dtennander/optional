package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalImpl_Map(t *testing.T) {
	// Given
	optA := Of(1)
	// When
	optB := optA.Map(func(a int) int {
		return a + 1
	})
	// Then
	assert.Equal(t, 2, optB.Get())
}

func TestOptionalImpl_Map2(t *testing.T) {
	// Given
	optA := Of(1)
	evilMap := func(a string) string { return a }
	// When
	evilCall := func() {
		optA.Map(evilMap)
	}
	// Then
	assert.PanicsWithValue(t, "Map function's argument is not compatible with this optional.", evilCall)
}

func TestOptionalImpl_Map3(t *testing.T) {
	// Given
	a := Of(1)
	// When
	evilCall := func() {
		a.Map(1) // Not a function
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a mapping function.", evilCall)
}

func TestMapOnEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	b := a.Map(func(a int) int {
		return a
	})
	// Then
	assert.False(t, b.IsPresent())
}
