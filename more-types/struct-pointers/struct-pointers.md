# Struct pointers

Struct fields can be accessed through struct pointers

In Go, there's a simple way of doing the following:
```go
p := &Vertex{1, 2}

(*p).X // Hard

p.X // Easy
```
