package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalImpl_Filter(t *testing.T) {
	// Given
	a := Of(1)
	// When
	b := a.Filter(func(i int) bool {
		return i == 1
	})
	// Then
	assert.True(t, b.IsPresent())
}

func TestOptionalImpl_Filter2(t *testing.T) {
	// Given
	a := Of(1)
	// When
	b := a.Filter(func(i int) bool {
		return i > 1
	})
	// Then
	assert.False(t, b.IsPresent())
}

func TestOptionalImpl_Filter3(t *testing.T) {
	// Given
	optA := Of(1)
	evilPredicate := func(i string) bool { return true }
	// When
	evilCall := func() {
		optA.Filter(evilPredicate)
	}
	// Then
	assert.PanicsWithValue(t, "Predicate's argument is not compatible with this optional.", evilCall)
}

func TestOptionalImpl_Filter4(t *testing.T) {
	// Given
	a := Of(1)
	evilPredicate := func(a int) int { return a }
	// When
	evilCall := func() {
		a.Filter(evilPredicate) // Not a function
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a predicate function.", evilCall)
}

func TestFilterOnEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	b := a.Filter(func(a int) bool {
		return true
	})
	// Then
	assert.False(t, b.IsPresent())
}
