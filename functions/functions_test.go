package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidConsumer(t *testing.T) {
	// Given
	consumer := func(i int) {}
	// When
	result := IsValid(Consumer, consumer)
	// Then
	assert.Equal(t, true, result)
}

func TestIsInvalidConsumer(t *testing.T) {
	// Given
	consumer := func() {}
	// When
	result := IsValid(Consumer, consumer)
	// Then
	assert.Equal(t, false, result)
}

func TestIsValidPredicate(t *testing.T) {
	// Given
	consumer := func(i int) bool { return true }
	// When
	result := IsValid(Predicate, consumer)
	// Then
	assert.Equal(t, true, result)
}

func TestIsInvalidPredicate(t *testing.T) {
	// Given
	consumer := func(i int) {}
	// When
	result := IsValid(Predicate, consumer)
	// Then
	assert.Equal(t, false, result)
}

func TestIsValidMap(t *testing.T) {
	// Given
	consumer := func(i int) int { return i + 1 }
	// When
	result := IsValid(Map, consumer)
	// Then
	assert.Equal(t, true, result)
}

func TestIsInvalidMap(t *testing.T) {
	// Given
	consumer := func(i int) {}
	// When
	result := IsValid(Map, consumer)
	// Then
	assert.Equal(t, false, result)
}

func TestTakesArgument(t *testing.T) {
	// Given
	consumer := func(i int) {}
	// When
	result := TakesArgument(consumer, 1)
	// Then
	assert.True(t, result)
}

func TestTakesArgument2(t *testing.T) {
	// Given
	consumer := func(i string) {}
	// When
	result := TakesArgument(consumer, 1)
	// Then
	assert.False(t, result)
}
