[![GoDoc](https://godoc.org/github.com/changkun/gomat?status.svg)](https://godoc.org/github.com/changkun/gomat)
[![Build Status](https://travis-ci.org/changkun/gomat.svg?branch=master)](https://travis-ci.org/changkun/gomat)
[![Go Report Card](https://goreportcard.com/badge/github.com/changkun/gomat)](https://goreportcard.com/report/github.com/changkun/gomat)

# gomat

Matrix package with cache-aware lock-free tiling optimization.

## Getting started

The following illustrates some basic usage of `gomat`.

```go
// Create 2x3 matrix, specified its value
// New() will throw error if provided values 
// is ineuqal to its dimension
A, err := gomat.New(2, 3)(
    1, 2, 3,
    4, 5, 6,
)

// Create a 3x4 random matrix
B := gomat.Rand(3, 4)

// Create a 2x4 zero matrix
C := gomat.Zero(2, 4)


// C = A x B, throw err if dimentionality error
err = A.DotNaive(B, C)
err = A.DotNaiveP(B, C) // with concurrency optimization
```

## License

MIT &copy; [changkun](https://changkun.de)