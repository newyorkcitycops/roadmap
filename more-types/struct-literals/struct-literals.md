# Struct literals

There's two forms of declaring structs

## Positional arguments
```go
v1 := Vertex{1, 2}
```

## Named arguments
```go
v2 := Vertex{X: 1}
```

Declaring a struct literal as a pointer, it's also possible
```go
p := &Vertex{1, 2}
```
