package fauxgaux

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestIntMap(t *testing.T) {
	things := NewChainable([]int{1, 2, 3, 4})
	things.Map(func(i int) int {
		return i + 1
	})
	fmt.Println(*things)
	fmt.Println(reflect.TypeOf(things))
}

func TestStringMap(t *testing.T) {
	words := NewChainable([]string{"hello", "what's up", "howdy"})
	words.Map(func(s string) string {
		return strings.Join([]string{s, "World!"}, " ")
	})
	fmt.Println(*words)
}
