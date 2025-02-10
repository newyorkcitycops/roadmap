# Images

image package defines the Image interface

An image interface defines ColorModel, Bounds and At as its methods signatures

```go
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
```

color.Color and color.Model are also interfaces, its concrete types can be color.RGBA and color.RGBAModel
