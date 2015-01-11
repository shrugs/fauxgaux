
# FauxGaux - Functional Go

Named FauxGaux because this isn't very idiomatic.

## Disclaimer

I have no idea what I'm doing with Golang _or_ functional programming and it's 5am.

## Usage

### Map

```go
things := Chain([]int{1, 2, 3, 4}).Map(func(i int) int {
    return i + 1
}).ConvertInt()
fmt.Println(things)
// [2 3 4 5]
```

```go
words := Chain([]string{"Hello", "What's up", "Howdy"}).Map(func(s string) string {
    return strings.Join([]string{s, "World!"}, " ")
}).ConvertString()
fmt.Println(words)
// [Hello World! What's up World! Howdy World!]
```

### Reduce

This section is currently under construction. Check back soon!