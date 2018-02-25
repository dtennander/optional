// Package Optional is a library constructed to allow for optional return values.
// Thus allowing the end user of APIs to get a nicer syntax in the end code.
// Instead of
//   val, err := vendor.getThing()
//   if err != nil {
//       return nil
//   }
//   return doStuff(val)
// One could write code:
//   vendor.getThing().Map(doStuff).orElse(nil)
package optional

import "reflect"

// The Optional type represents a possible value that might be there.
// It offers a API that allows for manipulating the possible value returning defaults.
type Optional interface {
	// Returns true if the Optional contains a value.
	IsPresent() bool

	// Returns the value contained in the Optional.
	// Will panic if called on an empty Optional.
	Get() interface{}

	// If a value is present the mapFunc is applied to the contained value and
	// the result is returned as an optional.
	// If called on an empty Optional an Empty optional is returned.
	Map(mapFunc interface{}) Optional

	// Returns an empty Optional if the predicate is not met.
	// Otherwise the current value is passed along.
	Filter(predicate interface{}) Optional

	// Returns the value if a value is present else it returns the defaultValue.
	OrElse(defaultValue interface{}) interface{}
}

// Of constructs an Optional from a value separate from nil.
func Of(nonNilValue interface{}) Optional {
	assertNotNil(nonNilValue)
	return &optionalImpl{isPresent: true, value: nonNilValue}
}

func assertNotNil(nonNilValue interface{}) {
	if isNil(nonNilValue) {
		panic("nonNilValue can not be nil!")
	}
}

func isNil(possibleNil interface{}) bool {
	value := reflect.ValueOf(possibleNil)
	return value.Kind() == reflect.Ptr && value.IsNil()
}

// OfPossibleNil construct an Optional from a given value.
// If the value is nil a empty Optional is returned.
func OfPossibleNil(value interface{}) Optional {
	if isNil(value) {
		return Empty()
	}
	return Of(value)

}

// Empty constructs an empty optional.
func Empty() Optional {
	return &optionalImpl{isPresent: false, value: nil}
}

type optionalImpl struct {
	isPresent bool
	value     interface{}
}

func (o *optionalImpl) Get() interface{} {
	if !o.isPresent {
		panic("Accessing empty Optional.")
	}
	return o.value
}

func (o *optionalImpl) OrElse(defaultValue interface{}) interface{} {
	if o.isPresent {
		return o.value
	}
	return defaultValue
}

func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}
