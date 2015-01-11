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

// same as Map but doesn't replace values in the array; only mofifies them in memory
// works with pointers to structs and stuff
func (c *Chainable) Each(fin interface{}) *Chainable {
	f := reflect.ValueOf(fin)
	t := f.Type()
	inType := t.In(0)
	for _, thing := range *c {

		v := reflect.New(inType).Elem()
		v.Set(reflect.ValueOf(thing))
		f.Call([]reflect.Value{v})

	}
	return c
}

// function passed to map takes function and pointer to accumulator
func (c *Chainable) Reduce(fin interface{}, accumulator interface{}) interface{} {
	f := reflect.ValueOf(fin)
	t := f.Type()
	objType := t.In(0)
	accType := t.In(1)

	a := reflect.New(accType).Elem()
	a.Set(reflect.ValueOf(accumulator))

	for _, thing := range *c {
		v := reflect.New(objType).Elem()
		v.Set(reflect.ValueOf(thing))
		out := f.Call([]reflect.Value{a, v})
		// set accumulator to whatever f returns
		a.Set(out[0])
	}

	return a.Interface()
}

// filter function returns a bool indicating whether or not to keep an element
func (c *Chainable) Filter(fin interface{}) *Chainable {
	newChainable := make(Chainable, 0, len(*c))

	f := reflect.ValueOf(fin)
	t := f.Type()
	inType := t.In(0)
	for _, thing := range *c {

		v := reflect.New(inType).Elem()
		v.Set(reflect.ValueOf(thing))
		out := f.Call([]reflect.Value{v})[0]
		if out.Bool() {
			newChainable = append(newChainable, thing)
		}
	}

	return &newChainable

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

// turns an array of objects into a chainable ([]interface{})
// takes a pointer to a slice of things
func Chain(in interface{}) *Chainable {
	t := reflect.TypeOf(in)
	v := reflect.New(t).Elem()
	v.Set(reflect.ValueOf(in))

	original := reflect.Indirect(v)

	newChainable := make(Chainable, original.Len(), original.Cap())

	for i := 0; i < original.Len(); i += 1 {
		newChainable[i] = original.Index(i).Interface()
	}
	return &newChainable
}
