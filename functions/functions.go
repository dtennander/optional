// Package functions contains functions and definitions useful for handling functions cast to interface{}.
package functions

import (
	"reflect"
)

type functionType int

const (
	// Predicate is any function taking one argument and returning a bool.
	//   f: X -> bool
	Predicate functionType = iota

	// Consumer is any function taking one argument and returning none.
	// Use this to produce side effects.
	Consumer

	// Map is any function taking one argument and returning another.
	//   f: X -> Y
	Map
)

// CallFunction calls the given function with the given argument and returns the result.
// Will panic if it is not possible.
func CallFunction(function interface{}, argument interface{}) reflect.Value {
	valValue := reflect.ValueOf(argument)
	funcValue := reflect.ValueOf(function)
	return funcValue.Call([]reflect.Value{valValue})[0]
}

// Consume calls the given function with the given argument.
// Will panic if it is not possible to do.
func Consume(consumer interface{}, argument interface{}) {
	valValue := reflect.ValueOf(argument)
	reflect.ValueOf(consumer).Call([]reflect.Value{valValue})
}

// TakesArgument returns true if the given function can take the given argument.
func TakesArgument(function interface{}, argument interface{}) bool {
	functionType := reflect.TypeOf(function)
	argumentType := reflect.TypeOf(argument)
	return isFunction(functionType) && argumentType.AssignableTo(functionType.In(0))
}

// IsValid returns true if the given function follows the rules for the given FunctionType.
func IsValid(functionType functionType, function interface{}) bool {
	funcType := reflect.ValueOf(function).Type()
	switch functionType {
	case Predicate:
		return isPredicate(funcType)
	case Consumer:
		return isConsumer(funcType)
	case Map:
		return isMap(funcType)
	default:
		return false
	}
}

func isPredicate(funcType reflect.Type) bool {
	return isFunction(funcType) &&
		isOneToOneFunction(funcType) &&
		funcType.Out(0).Kind() == reflect.Bool
}

func isConsumer(consumer reflect.Type) bool {
	return isFunction(consumer) && consumer.NumIn() == 1 && consumer.NumOut() == 0
}

func isMap(funcType reflect.Type) bool {
	return isFunction(funcType) && isOneToOneFunction(funcType)
}

func isFunction(funcType reflect.Type) bool {
	a := funcType.Kind() == reflect.Func
	return a
}

func isOneToOneFunction(funcType reflect.Type) bool {
	return funcType.NumIn() == 1 && funcType.NumOut() == 1
}
