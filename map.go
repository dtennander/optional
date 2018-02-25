package optional

import "reflect"

func (o *optionalImpl) Map(mapFunc interface{}) Optional {
	funcValue := reflect.ValueOf(mapFunc)
	checkIfTypeIsMap(funcValue.Type())
	if o.isPresent {
		return mapValue(o, funcValue)
	}
	return Empty()
}

func checkIfTypeIsMap(funcType reflect.Type) {
	if isNotFunction(funcType) || isNotOneToOneFunction(funcType) {
		panic("Argument must be a mapping function.")
	}
}

func isNotFunction(funcType reflect.Type) bool {
	return !isFunction(funcType)
}

func isFunction(funcType reflect.Type) bool {
	return funcType.Kind() == reflect.Func
}

func isNotOneToOneFunction(funcType reflect.Type) bool {
	return funcType.NumIn() != 1 || funcType.NumOut() != 1
}

func mapValue(o *optionalImpl, funcValue reflect.Value) Optional {
	valValue := reflect.ValueOf(o.value)
	valType := valValue.Type()
	funcType := funcValue.Type()
	// Checking whether element type is convertible to function's first argument's type.
	if !valType.AssignableTo(funcType.In(0)) {
		panic("Map function's argument is not compatible with this optional.")
	}
	result := funcValue.Call([]reflect.Value{valValue})[0]
	// Convert resulting slice back to generic interface.
	return OfPossibleNil(result.Interface())
}
