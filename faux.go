package fauxgaux

import (
	"fmt"
	"reflect"
)

type Chainable []interface{}

// function passed to map must return one and only one element of any type
func (c *Chainable) Map(fin interface{}) *Chainable {
	f := reflect.ValueOf(fin)
	t := f.Type()
	inType := t.In(0)
	for i, thing := range *c {

		v := reflect.New(inType).Elem()
		v.Set(reflect.ValueOf(thing))
		out := f.Call([]reflect.Value{v})
		(*c)[i] = out[0].Interface()
	}
	return c
}

// converts the interfaces to type defined by t
func (c *Chainable) Type(outType interface{}) []reflect.Value {

	t := reflect.ValueOf(outType)
	fmt.Println(t)

	// for _, thing := range c {

	// }
	return nil
}

// turns an array of ints into a Chainable
func NewChainable(in interface{}) *Chainable {
	original := reflect.ValueOf(in)

	copy := make(Chainable, original.Len(), original.Cap())

	for i := 0; i < original.Len(); i += 1 {
		copy[i] = original.Index(i).Interface()
	}

	return &copy
}
