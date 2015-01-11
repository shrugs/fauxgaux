package fauxgaux

import (
	"fmt"
	"strings"
	"testing"
)

func TestIntMap(t *testing.T) {
	things := Chain([]int{1, 2, 3, 4}).Map(func(i int) int {
		return i + 1
	}).ConvertInt()
	fmt.Println(things[0] + things[1])
	fmt.Println(things)
}

func TestStringMap(t *testing.T) {
	words := Chain([]string{"hello", "what's up", "howdy"}).Map(func(s string) string {
		return strings.Join([]string{s, "World!"}, " ")
	}).ConvertString()
	fmt.Println(words)
}
