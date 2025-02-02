# Type conversions

Integers can be converted to floating-point numbers (floats)

Booleans can be converted to numbers

In Go, there is no such package or name for type conversion, type call itself can convert other types

```go
var i int = 42
var f float64 = float64(i)
```

Whatsoever, string conversions are handled by strconv package

```go
// Incorrect:
i := int("42")

// Correct:
i := strconv.Atoi("42")
s := strconv.Itoa(42)
```
