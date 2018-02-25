package optional

import "reflect"

func (o *optionalImpl) Filter(predicateFunc interface{}) Optional {
	funcValue := reflect.ValueOf(predicateFunc)
	checkIfTypeIsPredicate(funcValue.Type())
	if o.isPresent {
		return applyPredicate(o, funcValue)
	} else {
		return Empty()
	}
}
func applyPredicate(o *optionalImpl, funcValue reflect.Value) Optional {
	valValue := reflect.ValueOf(o.value)
	valType := valValue.Type()
	funcType := funcValue.Type()
	// Checking whether element type is convertible to function's first argument's type.
	if !valType.ConvertibleTo(funcType.In(0)) {
		panic("Predicate's argument is not compatible with this optional.")
	}

	predicateIsTrue := funcValue.Call([]reflect.Value{valValue})[0].Bool()
	if predicateIsTrue {
		return o
	} else {
		return Empty()
	}
}

func checkIfTypeIsPredicate(funcType reflect.Type) {
	if isNotFunction(funcType) || isNotOneToOneFunction(funcType) || isNotPredicate(funcType) {
		panic("Argument must be a predicate function.")
	}
}

func isNotPredicate(funcType reflect.Type) bool {
	return funcType.Out(0).Kind() != reflect.Bool
}
