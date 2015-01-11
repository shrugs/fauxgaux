package fauxgaux

import (
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

func (c *Chainable) ConvertInt() []int {
	convertedSlice := make([]int, len(*c))

	for i, t := range *c {
		convertedSlice[i] = t.(int)
	}

	return convertedSlice
}

func (c *Chainable) ConvertString() []string {
	convertedSlice := make([]string, len(*c))

	for i, t := range *c {
		convertedSlice[i] = t.(string)
	}

	return convertedSlice
}

// turns an array of ints into a Chainable
func Chain(in interface{}) *Chainable {
	original := reflect.ValueOf(in)

	newChainable := make(Chainable, original.Len(), original.Cap())

	for i := 0; i < original.Len(); i += 1 {
		newChainable[i] = original.Index(i).Interface()
	}
	return &newChainable
}
