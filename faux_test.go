package fauxgaux

import (
	"math"
	"strings"
	"testing"
)

func expect(t *testing.T, a, b interface{}) {
	if !(a == b) {
		t.Errorf("Expected: %v, Got: %v", a, b)
	}
}

type Person struct {
	Name string
	Age  int
}

func TestIntMap(t *testing.T) {
	nums := Chain(&[]int{1, 2, 3, 4}).Map(func(i int) int {
		return i + 1
	}).ConvertInt()
	expect(t, nums[0]+nums[1], 5) // expect 2 + 3 == 5
}

func TestStringMap(t *testing.T) {
	words := Chain(&[]string{"Hello", "What's up", "Howdy"}).Map(func(s string) string {
		return strings.Join([]string{s, "World!"}, " ")
	}).ConvertString()
	expect(t, words[0], "Hello World!")
}

func TestReduceInt(t *testing.T) {
	sum := Chain(&[]int{1, 2, 3, 4, 5}).Reduce(func(i int, num int) int {
		i += num
		return i
	}, 0).(int)
	expect(t, sum, 15)
}

func TestMapReduceStruct(t *testing.T) {
	people := &[]*Person{
		&Person{"Matt", 20},
		&Person{"Ben", 19},
	}

	totalAge := Chain(people).Map(func(p *Person) int {
		return p.Age
	}).Reduce(func(i int, num int) int {
		i += num
		return i
	}, 0).(int)

	expect(t, totalAge, 39)
}

func TestMapReduceModifyStruct(t *testing.T) {
	people := &[]*Person{
		&Person{"Matt", 20},
		&Person{"Ben", 19},
		&Person{"Sam", 21},
	}

	Chain(people).Each(func(p *Person) {
		p.Name = "test"
	})

	for _, p := range *people {
		expect(t, p.Name, "test")
	}

}

func TestFilterEven(t *testing.T) {
	nums := &[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSum := Chain(nums).Filter(func(i int) bool {
		return math.Mod(float64(i), 2) == 0
	}).Reduce(func(sum, num int) int {
		sum += num
		return sum
	}, 0).(int)

	expect(t, evenSum, 30)

}
