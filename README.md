
# FauxGaux - Functional Go

Named FauxGaux because this isn't very idiomatic.

## Disclaimer

This is for learning purposes and the hell of it. _Please_ do not actually use in a project.

This probably has an awful memory profile; don't look at it.

I have given minimal thought to error handling. As such, there is none.

## Installation

`go get github.com/Shrugs/fauxgaux`

## Usage

### Map

Takes a function. This function must accept one argument of arbitrary type and return one argument of arbitrary type.

```go
things := fauxgaux.Chain([]int{1, 2, 3, 4}).Map(func(i int) int {
    return i + 1
}).ConvertInt()
fmt.Println(things)
// [2 3 4 5]
```

```go
words := fauxgaux.Chain([]string{"Hello", "What's up", "Howdy"}).Map(func(s string) string {
    return strings.Join([]string{s, "World!"}, " ")
}).ConvertString()
fmt.Println(words)
// [Hello World! What's up World! Howdy World!]
```

### Reduce

Takes a function and an accumulator. Function must accept the accumulator and the current object (using the correct types, naturally) and return the accumulator (with type unchanged).

User must instantiate accumulator to some value.

```go
people := &[]*Person{
    &Person{"Matt", 20},
    &Person{"Ben", 19},
}

totalAge := fauxgaux.Chain(people).Map(func(p *Person) int {
    return p.Age
}).Reduce(func(i int, num int) int {
    i += num
    return i
}, 0).(int)

fmt.Println(totalAge)
// 39
```

### Each

Same as Map, but does not replace the values of the chained slice; it only modifies them in memory. As such, the function passed must only accept one argument of the correct type.

```go
type Person struct {
    Name string
    Age  int
}

people := &[]*Person{
    &Person{"Matt", 20},
    &Person{"Ben", 19},
    &Person{"Sam", 21},
}

fauxgaux.Chain(people).Each(func(p *Person) {
    p.Name = "test"
})

for _, p := range *people {
    fmt.Println(p.Name)
}

// test test test
```

### Filter

Filters the slice based on a conditional function. Function should accept one argument of the correct type and return a bool indicating whether or not to keep the element in the slice.

```go
nums := &[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenSum := fauxgaux.Chain(nums).Filter(func(i int) bool {
    return math.Mod(float64(i), 2) == 0
}).Reduce(func(sum, num int) int {
    sum += num
    return sum
}, 0).(int)

fmt.Println(evenSum)
// 30
```