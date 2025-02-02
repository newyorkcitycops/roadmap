# Named return values

They are treated as variables, so it means you can assign their values

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

A "naked" return can take place after