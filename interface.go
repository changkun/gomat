package gomat

// Matrix interface
type Matrix interface {
	Row() int
	Col() int
	Size() (int, int)
	At(i, j int) float64
	Set(i, j int, val float64)
	Inc(i, j int, val float64)
	Mult(i, j int, val float64)
	Pow(i, j int, n float64)
}
