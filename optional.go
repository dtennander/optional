package optional

type Optional interface {
	IsPresent() bool
	Get() interface{}
	Map(mapFunc interface{}) Optional
	Filter(predicate interface{}) Optional
}

type optionalImpl struct {
	isPresent bool
	value interface{}
}

func (o *optionalImpl) Get() interface{} {
	return o.value
}

func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}

func Of(value interface {}) Optional {
	return &optionalImpl{isPresent:true, value:value}
}

func Empty() Optional {
	return &optionalImpl{isPresent:false, value:nil}
}






