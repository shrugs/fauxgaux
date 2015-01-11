
# FauxGaux - Functional Go

Named FauxGaux because this isn't very idiomatic.

## Disclaimer

I have no idea what I'm doing with Golang _or_ functional programming and it's 5am.

## Installation

`go get github.com/Shrugs/fauxgaux`

## Usage

### Map

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

```go
people := &[]Person{
    Person{"Matt", 20},
    Person{"Ben", 19},
}

totalAge := Chain(people).Map(func(p Person) int {
    return p.Age
}).Reduce(func(i int, num int) int {
    i += num
    return i
}, 0).(int)

fmt.Println(totalAge)
// 39
```

### Each

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